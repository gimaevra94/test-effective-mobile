// Пакет предоставляет модуль для взаимодействия с базой данных.
package database

import (
	"database/sql"

	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	InsertQuery = "insert into subscriptions (service_name, price, user_id, start_date) values ($1, $2, $3, $4)"
)

type DB struct {
	*sql.DB
}

// Функция открывает соединение с базой данных.
func DBConn(cfg string) (*DB, error) {
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}
	return &DB{db}, nil
}

// Функция реализует операцию "create" добавляя в базу данных новую строку.
// Проверяет наличие дубля и в случае его отсутствия добавляет поля структуры базу данных.
func (db *DB) CreateSubscription(sub *structs.Subscription) error {
	if _, err := db.DB.Exec(InsertQuery, sub.ServiceName, sub.Price, sub.UserId, sub.StartDate); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			err := errors.New("Already exists")
			return errors.WithStack(err)
		}
		return errors.WithStack(err)
	}
	return nil
}
