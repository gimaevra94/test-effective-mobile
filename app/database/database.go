package database

import (
	"database/sql"

	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/pkg/errors"
)

const (
	InsertQuery = "insert into subscriptions (service_name, price, user_id, start_date) values (?,?,?,?)"
)

type DB struct {
	*sql.DB
}

func DBConn() (*DB, error) {
	var cfg string
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}

	return &DB{db}, nil
}

func (db *DB) CreateSubscription(sub *structs.Subscription) error {
	if _, err := db.Exec(InsertQuery, sub.ServiceName, sub.Price, sub.UserId, sub.StartDate); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
