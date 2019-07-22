package main

import (
	"./api"
	"./db"
	"./search"
)

func main() {
	storeDB := db.InitDB()
	defer storeDB.DB.Close()

	searchDB := search.InitSearch()

	searchDB.CreateIndexFromBadgerDB(storeDB)
	// store.DownloadProducts(storeDB)
	// store.ReadAllProducts()
	api.Listen(storeDB, searchDB)
}
