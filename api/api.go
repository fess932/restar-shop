package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"restar-shop/db"
	"restar-shop/search"

	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
)

// JSON s
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

const localhost = "localhost:8080"

// Listen открываем порт на сервере
func Listen(storeDB *db.Store, searchDB *search.DB) {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("root"))
		if err != nil {
			log.Println(err)
		}
	})

	r.Get("/search/{searchString}", func(w http.ResponseWriter, r *http.Request) {
		searchString := chi.URLParam(r, "searchString")
		qs := normalizeQuery(searchString)

		items := searchDB.Search(qs)

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

func normalizeQuery(s string) (qs string) {
	s = strings.ToUpper(s)
	s = Replacer(s)
	qs = s
	return
}
