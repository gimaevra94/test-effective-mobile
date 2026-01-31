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
	if err := database.DBConn(); err != nil {
		log.Fatalln(err)
	}
	defer database.DB.Close()

	r := initRouter()
	http.ListenAndServe(":8080", r)

}

func initRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/v1/subscription", handlers.CreateSubscription)
	return r
}
