package main

import (
	"./db"
)

func main() {
	store := db.InitDB()
	defer store.DB.Close()

	store.DownloadProducts()
	// store.ReadAllProducts()
	// api.Listen()
}
