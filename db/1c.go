package db

import (
	"io/ioutil"
	"log"
	"restar-shop/utilits"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Product is product from 1c json
type Product struct {
	Name           string  `json:"Наименование"`
	Characteristic string  `json:"Характеристика"`
	FullName       string  `json:"НаименованиеПолное"`
	Code           string  `json:"Код"`
	GUID           string  `json:"GUID"`
	ParentName     string  `json:"НаименованиеРодителя"`
	ParentGUID     string  `json:"GUIDРодителя"`
	Manufacture    string  `json:"Производитель"`
	SKU            string  `json:"Артикул"`
	Properties     string  `json:"Габариты"`
	StockQuantity  int     `json:"Остаток"`
	Price          float32 `json:"Цена"`
}

func returnProducts() []Product {
	var products []Product

	// file, err := ioutil.ReadFile("data/nomenklatura.json")
	file, err := ioutil.ReadFile("data/dataFULL.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &products)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range products {
		products[i].Name = utilits.Replacer(v.Name)
		products[i].SKU = utilits.Replacer(v.SKU)
	}

	return products
}
