package main

import (
	"./api"
	"./product"
)

func main() {
	c := product.ReadProducts("data/data.json")
	api.Listen(c)
}
