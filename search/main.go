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

// Item Define struct with reindex tags
type Item struct {
	ID       int64  `json:"id" reindex:"id,,pk"`           // 'id' is primary key
	Name     string `json:"name" reindex:"name,fuzzytext"` // add index by 'name' field
	Articles []int  `json:"articles" reindex:"articles"`   // add index by articles 'articles' array
	Year     int    `json:"year" reindex:"year,tree"`      // add sortable index by 'year' field
}

// Shop is shop
type Shop struct {
	Items []Item
	db    *reindexer.Reindexer
}

func createIn(db *reindexer.Reindexer) {

	// Generate dataset
	for i := 0; i < 10; i++ {
		err := db.Upsert("items", &Item{
			ID:       int64(i),
			Name:     "Вася cs 1050",
			Articles: []int{rand.Int() % 100, rand.Int() % 100},
			Year:     2000 + rand.Int()%50,
		})
		if err != nil {
			panic(err)
		}
	}

	for i := 10; i < 20; i++ {
		err := db.Upsert("items", &Item{
			ID:       int64(i),
			Name:     "Петя cs.1050",
			Articles: []int{rand.Int() % 100, rand.Int() % 100},
			Year:     2000 + rand.Int()%50,
		})
		if err != nil {
			panic(err)
		}

		for i := 20; i < 30; i++ {
			err := db.Upsert("items", &Item{
				ID:       int64(i),
				Name:     "Вова cs1050",
				Articles: []int{rand.Int() % 100, rand.Int() % 100},
				Year:     2000 + rand.Int()%50,
			})
			if err != nil {
				panic(err)
			}
		}
	}

}

// InitSearch get instance of search
func InitSearch() *Shop {
	shop := Shop{}
	shop.db = reindexer.NewReindex("cproto://127.0.0.1:6534/testdb")
	shop.db.OpenNamespace("items", reindexer.DefaultNamespaceOptions(), Item{})

	return &shop
}

// Search searching items and get array of items if exist
func (shop *Shop) Search(qs string) []Item {
	items := []Item{}

	query := shop.db.Query("items").
		Match("Name", qs)

		// Execute the query and return an iterator
	iterator := query.Exec()
	// Iterator must be closed
	defer iterator.Close()

	fmt.Println("Found", iterator.TotalCount(), "total documents, first", iterator.Count(), "documents:")

	// Iterate over results
	for iterator.Next() {
		// Get the next document and cast it to a pointer
		elem := iterator.Object().(*Item)
		items = append(items, *elem)
		// fmt.Println(*elem)
	}
	// Check the error
	if err := iterator.Error(); err != nil {
		panic(err)
	}

	return items

}
