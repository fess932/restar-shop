package product

import (
	"fmt"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

type Product struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	SKU  string `json:"sku"`
}

type Config struct {
	Products map[string]Product  `json:"products"`
	SKU      map[string][]string `json:"sku"`
}

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
