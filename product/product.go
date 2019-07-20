package product

import (
	"fmt"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

// Product is one product
type Product struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	SKU  string `json:"sku"`
}

// Config is full config of db
type Config struct {
	Products []Product `json:"products"`
}

// ReadProducts reding products
func ReadProducts(path string) *Config {
	c := Config{}

	fmt.Println("reading products")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	Json.Unmarshal(file, &c)
	return &c
}
