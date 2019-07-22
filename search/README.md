# reindexer

// установить модуль reindexer

// запустить докер с сервером reindexer
docker run -p9088:9088 -p6534:6534 -it reindexer/reindexer

## Для нечеткого поиска с индексом по триграммам использовать ключевое слово "fuzzytext"

```lang=golang
package main

// Import package
import (
  "fmt"
  "math/rand"

  "github.com/restream/reindexer"
  // choose how the Reindexer binds to the app (in this case "builtin," which means link Reindexer as a static library)
  _ "github.com/restream/reindexer/bindings/builtin"

  // OR link Reindexer as static library with bundled server.
  // _ "github.com/restream/reindexer/bindings/builtinserver"
  // "github.com/restream/reindexer/bindings/builtinserver/config"

)

// Define struct with reindex tags
type Item struct {w
  ID       int64  `reindex:"id,,pk"`    // 'id' is primary key
  Name     string `reindex:"name,fuzzytext"`      // add index by 'name' field
  Articles []int  `reindex:"articles"`  // add index by articles 'articles' array
  Year     int    `reindex:"year,tree"` // add sortable index by 'year' field
}

func main() {
  // Init a database instance and choose the binding (builtin)
  db := reindexer.NewReindex("builtin:///tmp/reindex/testdb")

  // OR - Init a database instance and choose the binding (connect to server)
  // db := reindexer.NewReindex("cproto://127.0.0.1:6534/testdb")

  // OR - Init a database instance and choose the binding (builtin, with bundled server)
  // serverConfig := config.DefaultServerConfig ()
  // db := reindexer.NewReindex("builtinserver://testdb",reindexer.WithServerConfig(100*time.Second, serverConfig))

  // Create new namespace with name 'items', which will store structs of type 'Item'
  db.OpenNamespace("items", reindexer.DefaultNamespaceOptions(), Item{})

  // Generate dataset
  for i := 0; i < 100000; i++ {
    err := db.Upsert("items", &Item{
      ID:       int64(i),
      Name:     "Vasya",
      Articles: []int{rand.Int() % 100, rand.Int() % 100},
      Year:     2000 + rand.Int()%50,
    })
    if err != nil {
      panic(err)
    }
  }

  // Query a single document
  elem, found := db.Query("items").
    Where("id", reindexer.EQ, 40).
    Get()

  if found {
    item := elem.(*Item)
    fmt.Println("Found document:", *item)
  }

  // Query multiple documents
  query := db.Query("items").
    Sort("year", false).                          // Sort results by 'year' field in ascending order
    WhereString("name", reindexer.EQ, "Vasya").   // 'name' must be 'Vasya'
    WhereInt("year", reindexer.GT, 2020).         // 'year' must be greater than 2020
    WhereInt("articles", reindexer.SET, 6, 1, 8). // 'articles' must contain one of [6,1,8]
    Limit(10).                                    // Return maximum 10 documents
    Offset(0).                                    // from 0 position
    ReqTotal()                                    // Calculate the total count of matching documents

  // Execute the query and return an iterator
  iterator := query.Exec()
  // Iterator must be closed
  defer iterator.Close()

  fmt.Println("Found", iterator.TotalCount(), "total documents, first", iterator.Count(), "documents:")

  // Iterate over results
  for iterator.Next() {
    // Get the next document and cast it to a pointer
    elem := iterator.Object().(*Item)
    fmt.Println(*elem)
  }
  // Check the error
  if err := iterator.Error(); err != nil {
    panic(err)
  }
}
```
