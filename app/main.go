package main

import (
	"log"
	"net/http"

	"github.com/gimaevra94/test-effective-mobile/app/database"
	"github.com/gimaevra94/test-effective-mobile/app/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := database.DBConn(); err != nil {
		log.Fatalln(err)
	}

	r := initRouter()
	http.ListenAndServe(":8080", r)

	defer database.DB.Close()
}

func initRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("api/create-subscription", handlers.CreateSubscription)
	return r
}
