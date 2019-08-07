package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"restar-shop/db"
	"restar-shop/search"
	"restar-shop/utilits"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	jsoniter "github.com/json-iterator/go"
)

// JSON s
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

const localhost = "localhost:8080"

// Listen открываем порт на сервере
func Listen(storeDB *db.Store, searchDB *search.DB) {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

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

		fmt.Println(qs)
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
	s = utilits.Replacer(s)
	qs = s
	return
}
