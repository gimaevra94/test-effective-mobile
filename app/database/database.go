package database

import (
	"database/sql"

	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/pkg/errors"
)

const (
	InsertQuery = "insert into subscriptions (service_name, price, user_id, start_date) values (?,?,?,?) on duplicate key update service_name = values(service_name), price = values(price), user_id = values(user_id), start_date = values(start_date)"
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

	return nil
}

func CreateSubscription(sub *structs.Subscription) error {
	if _, err := DB.Exec(InsertQuery, sub.ServiceName, sub.Price, sub.UserId, sub.StartDate); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
