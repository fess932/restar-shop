package db

import (
	"io/ioutil"
	"log"

	"github.com/json-iterator/go"
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

	file, err := ioutil.ReadFile("data/nomenklatura.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &products)
	if err != nil {
		log.Fatal(err)
	}

	return products
}
