package search

// Import package
import (
	"fmt"
	"io/ioutil"
	"log"

	"../db"
	jsoniter "github.com/json-iterator/go"
	"github.com/restream/reindexer"

	// choose how the Reindexer binds to the app (in this case "builtin," which means link Reindexer as a static library)
	_ "github.com/restream/reindexer/bindings/builtin"
	// OR link Reindexer as static library with bundled server.
	// _ "github.com/restream/reindexer/bindings/builtinserver"
	// "github.com/restream/reindexer/bindings/builtinserver/config"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Item Define struct with reindex tags
type Item struct {
	ID   string `json:"id" reindex:"id,,pk"`                   // 'id' is primary key
	Name string `json:"Наименование" reindex:"name,fuzzytext"` // add index by 'name' field
	SKU  string `json:"sku" reindex:"sku,fuzzytext"`           // add sortable index by 'year' field
}

// Item Define struct with reindex tags
// type Item struct {
// 	ID   string `json:"id" reindex:"id,,pk"`      // 'id' is primary key
// 	Name string `json:"name" reindex:"name,text"` // add index by 'name' field
// 	SKU  string `json:"sku" reindex:"sku,text"`   // add sortable index by 'year' field
// }

// DB is search db
type DB struct {
	Items []Item `json:"items"`
	db    *reindexer.Reindexer
}

func (shop *DB) createIndex() {

	file, err := ioutil.ReadFile("data/data.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &shop.Items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(shop.Items)
	// Generate dataset
	for _, v := range shop.Items {
		err := shop.db.Upsert("items", &Item{
			ID:   v.ID,
			Name: v.Name,
			SKU:  v.SKU,
		})
		if err != nil {
			panic(err)
		}
	}
}

// CreateIndexFromBadgerDB create index from budger db
func (shop *DB) CreateIndexFromBadgerDB(db *db.Store) {

}

// Search searching items and get array of items if exist
func (shop *DB) Search(qs string) []Item {
	items := []Item{}

	query := shop.db.Query("items").
		Match("Name", qs)

	query2 := shop.db.Query("items").
		Match("sku", qs)

	query.Merge(query2)
	// Execute the query and return an iterator
	iterator := query.Exec()
	// Iterator must be closed
	defer iterator.Close()

	// Check the error
	if err := iterator.Error(); err != nil {
		panic(err)
	}

	fmt.Println("Found", iterator.TotalCount(), "total documents, first", iterator.Count(), "documents:")

	// Iterate over results
	for iterator.Next() {
		// Get the next document and cast it to a pointer
		elem := iterator.Object().(*Item)
		fmt.Println(elem)
		items = append(items, *elem)
		// fmt.Println(*elem)
	}
	// Check the error
	if err := iterator.Error(); err != nil {
		panic(err)
	}

	return items

}

// InitSearch get instance of search
func InitSearch() *DB {
	shop := DB{}
	shop.db = reindexer.NewReindex("cproto://127.0.0.1:6534/testdb")
	shop.db.OpenNamespace("items", reindexer.DefaultNamespaceOptions(), Item{})
	shop.createIndex()

	return &shop
}
