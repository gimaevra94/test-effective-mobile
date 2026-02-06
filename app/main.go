package main

import (
	"log"
	"net/http"

	"github.com/gimaevra94/test-effective-mobile/app/consts"
	"github.com/gimaevra94/test-effective-mobile/app/database"
	"github.com/gimaevra94/test-effective-mobile/app/handlers"
	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Открывает соединение с базой данных.
// Инициализирует роутер.
// Запускает сервер
func main() {
	db, gdb, err, gErr := initDB()
	if err != nil && gErr != nil {
		log.Fatalln(err, gErr)
	}

	r := initRouter(db, gdb)
	http.ListenAndServe(":8080", r)
}

// Открывает пул db для: create, read, update, delete,
// и gdb для list.
func initDB() (*database.DB, *gorm.DB, error, error) {
	var cfg string
	db, err := database.DBConn(cfg)
	if err != nil {
		return nil, nil, errors.WithStack(err), nil
	}
	defer db.DB.Close()

	gdb, err := gorm.Open(postgres.Open(cfg), &gorm.Config{})
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}
	gdb.AutoMigrate(&structs.Subscription{})

	sqlDB, err := gdb.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	return db, gdb, nil, nil
}

// Обработчик маршрутов
func initRouter(db *database.DB, gdb *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Post(consts.APIPathV1, handlers.CreateSubscription(db))
	r.Get(consts.APIPathV1+"/{"+consts.UserID+"}/{"+consts.ServiceName+"}", handlers.GetSubscription(db))
	r.Patch(consts.APIPathV1+"/{"+consts.UserID+"}/{"+consts.ServiceName+"}", handlers.UpdateSubscription(db))
	r.Delete(consts.APIPathV1, handlers.DeleteSubscription(db))
	r.Get(consts.APIPathV1, handlers.ListSubscription(gdb))

	return r
}
