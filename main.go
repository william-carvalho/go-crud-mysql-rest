package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/william-carvalho/go-crud-mysql-rest/driver"
	ph "github.com/william-carvalho/go-crud-mysql-rest/handler/http"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", postRouter(pHandler))
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)

}

func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Get("/", pHandler.Fetch)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)

	return r
}
