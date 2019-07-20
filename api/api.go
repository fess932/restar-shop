package api

import (
	"fmt"
	"log"
	"net/http"

	"../product"
	s "../search"

	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
)

// JSON s
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

const localhost = "localhost:8080"

// Listen открываем порт на сервере
func Listen(c *product.Config) {
	r := chi.NewRouter()
	search := s.InitSearch()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("root"))
		if err != nil {
			log.Println(err)
		}
	})

	r.Get("/search/{searchString}", func(w http.ResponseWriter, r *http.Request) {
		searchString := chi.URLParam(r, "searchString")

		items := search.Search(searchString)

		itemsMarshal, _ := JSON.Marshal(items)

		w.Write(itemsMarshal)

		fmt.Println(searchString)
	})

	go http.ListenAndServe(localhost, r)

	fmt.Println("api listen at", localhost)

	// ? scanning in terminal

	scancode()

	fmt.Println("done")
}

func scancode() {
	scancode := ""
	println(`input "stop" to stop`)
	for {
		if scancode == "stop" {
			break
		}
		_, err := fmt.Scanln(&scancode)
		if err != nil {
			log.Fatal(err)
		}
	}

}
