package main

import (
	"fmt"
	"restar-shop/db"
)

func main() {
	// init DB for shop

	fmt.Println("hello world")
	storeDB := db.InitDB()
	defer storeDB.DB.Close()

	// dowwload products from 1c json
	// storeDB.DownloadProducts()

	// init search indexer for shop
	// searchDB := search.InitSearch()
	// searchDB.CreateIndexFromBadgerDB(storeDB)

	//storeDB.ReadAllProducts()
	//api.Listen(storeDB, searchDB)
}
