package api

import (
	"fmt"
	"log"
	"net/http"
	"restarShopGo/product"

	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

const localhost = "localhost:8080"

func Listen(c *product.Config) {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("root"))
		if err != nil {
			log.Println(err)
		}
	})

	r.Get("/search/{searchString}", func(w http.ResponseWriter, r *http.Request) {
		searchString := chi.URLParam(r, "searchString")
		sliceIDS := SearchID(c, searchString)

		if len(sliceIDS) == 0 {
			_, err := w.Write([]byte("Ничего не найдено"))
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		sliceProducts := returnProductsByID(c, sliceIDS)

		data, _ := Json.Marshal(sliceProducts)
		w.Write(data)

		fmt.Println(searchString)
	})

	go http.ListenAndServe(localhost, r)

	fmt.Println("api listen at", localhost)

	// ? scanning in terminal
	scancode := ""
	for {
		if scancode == "stop" {
			break
		}
		_, err := fmt.Scanln(&scancode)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

func SearchID(c *product.Config, s string) []string {
	var sKeys []string

	if val := searchSKU(c, s); val != "" {
		for _, v := range c.SKU[val] {
			sKeys = append(sKeys, v)
		}
	}

	return sKeys
}

func searchSKU(c *product.Config, s string) string {
	if val, ok := c.Products[s]; ok {
		return val.SKU
	}
	return ""
}

func returnProductsByID(c *product.Config, IDS []string) []product.Product {
	var ps []product.Product
	for _, v := range IDS {
		ps = append(ps, c.Products[v])
	}
	return ps
}
