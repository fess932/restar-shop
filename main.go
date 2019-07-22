package main

import (
	"restar-shop/db"
)

func main() {
	storeDB := db.InitDB()
	defer storeDB.DB.Close()

	// searchDB := search.InitSearch()

	// searchDB.CreateIndexFromBadgerDB(storeDB)
	//storeDB.DownloadProducts()
	storeDB.ReadAllProducts()
	// api.Listen(storeDB, searchDB)
}
