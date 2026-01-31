package main

import (
	"log"
	"net/http"

	"github.com/gimaevra94/test-effective-mobile/app/database"
	"github.com/gimaevra94/test-effective-mobile/app/handlers"
	"github.com/go-chi/chi/v5"
)

// Открывает соединение с базой данных.
// Инициализирует роутер.
// Запускает сервер
func main() {
	var cfg string
	db, err := database.DBConn(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.DB.Close()

	r := initRouter(db)
	http.ListenAndServe(":8080", r)
}

func initRouter(db *database.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/v1/subscription", handlers.CreateSubscription(db))
	r.Get("/api/v1/subscription", handlers.GetSubscription(db))
	
	return r
}
