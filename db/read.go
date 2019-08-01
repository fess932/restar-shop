package db

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

// ReadAllProducts чтение всех продуктов из базы данных и показ их
func (s *Store) ReadAllProducts() [][]byte {

	var kv [][]byte

	err := s.DB.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("products")

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			//k := item.Key()
			var cp []byte
			cp, err := item.ValueCopy(cp)
			if err != nil {
				log.Fatal(err)
			}
			kv = append(kv, cp)
			//err := item.Value(func(v []byte) error {
			//
			//	if string(k) == "products01b6ac76-4d63-11e5-949b-08606ed666b8" {
			//		fmt.Println(string(v))
			//	}
			//
			//	kv = append(kv, v)
			//	return nil
			//})
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	})

	fmt.Println("длинна массива значений: ", len(kv))
	if err != nil {
		log.Fatal(err)
	}

	return kv
}
