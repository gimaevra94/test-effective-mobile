package database

import (
	"database/sql"

	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/pkg/errors"
)

const (

	uniqueUserService = "create uniaue index if not exists unique_user_service on subscriptions (user_id, service_name);"
	InsertQuery = "insert into subscriptions (service_name, price, user_id, start_date) values ($1, $2, $3, $4) on conflict (user_id, service_name) do update est price = excluded.price"
)

var DB *sql.DB

func DBConn() error {
	var cfg string
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := db.Ping(); err != nil {
		return errors.WithStack(err)
	}

	DB=db
	return nil
}

func CreateSubscription(sub *structs.Subscription) error {
	if _, err := DB.Exec(InsertQuery, sub.ServiceName, sub.Price, sub.UserId, sub.StartDate); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
