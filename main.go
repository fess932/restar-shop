package main

import (
	"restarShopGo/api"
	"restarShopGo/product"
)

func main() {
	c := product.ReadProducts("data/data.json")
	api.Listen(c)
}
