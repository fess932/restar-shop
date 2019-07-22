package db

import (
	"log"

	"github.com/dgraph-io/badger"
)

// Store is main type for db and ever
type Store struct {
	DB *badger.DB
}

// InitDB initialize db
func InitDB() *Store {
	var store Store
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))
	store.DB = db
	if err != nil {
		log.Fatal(err)
	}
	return &store
}

// DownloadProducts get products from json or anythig
func (s *Store) DownloadProducts() {
	prefix := "products"
	products := returnProducts()

	for _, v := range products {
		key := prefix + v.GUID + v.Characteristic
		value, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}

		err = s.DB.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte(key), []byte(value))
			return err
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
