package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gimaevra94/test-effective-mobile/app/consts"
	"github.com/gimaevra94/test-effective-mobile/app/database"
	"github.com/gimaevra94/test-effective-mobile/app/handlers"
	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// @title Subscription API
// @version 1.0
// @description API для управления подписками пользователей на сервисы
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// Открывает соединение с базой данных.
// Инициализирует роутер.
// Запускает сервер
func main() {
	initEnv()

	db, gdb, err, gErr := initDB()
	if err != nil || gErr != nil {
		log.Fatal(err, gErr)
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	defer db.DB.Close()

	r := initRouter(db, gdb)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func initEnv() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
		return
	}

	envVars := []string{
		"CONNECTION_CFG",
		"POSTGRES_PASSWORD",
	}

	for _, v := range envVars {
		if os.Getenv(v) == "" {
			log.Fatal(v)
			return
		}
	}
}

// Открывает пул db для: create, read, update, delete,
// и gdb для list.
func initDB() (*database.DB, *gorm.DB, error, error) {
	cfg := os.Getenv("CONNECTION_CFG")
	db, err := database.DBConn(cfg)
	if err != nil {
		return nil, nil, errors.WithStack(err), nil
	}

	gdb, err := gorm.Open(postgres.Open(cfg), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, nil, nil, errors.WithStack(err)
	}

	if err := gdb.AutoMigrate(&structs.Subscription{}); err != nil {
		log.Fatal(err)
	}

	return db, gdb, nil, nil
}

// Обработчик маршрутов
func initRouter(db *database.DB, gdb *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Post(consts.APIPathV1, handlers.CreateSubscription(db))
	r.Get(consts.APIPathV1+"/{"+consts.ServiceName+"}/{"+consts.UserID+"}", handlers.GetSubscription(db))
	r.Patch(consts.APIPathV1+"/{"+consts.ServiceName+"}/{"+consts.UserID+"}", handlers.UpdateSubscription(db))
	r.Delete(consts.APIPathV1+"/{"+consts.ServiceName+"}/{"+consts.UserID+"}", handlers.DeleteSubscription(db))
	r.Get(consts.APIPathV1, handlers.ListSubscription(gdb))
	r.Get(consts.APIPathV1+"/totalPrice", handlers.GetPeriodTotalPrice(db))

	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
	))

	return r
}
