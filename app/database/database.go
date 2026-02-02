// Пакет предоставляет модуль для взаимодействия с базой данных.
package database

import (
	"database/sql"

	"github.com/gimaevra94/test-effective-mobile/app/consts"
	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type DB struct {
	*sql.DB
}

// Функция открывает соединение с базой данных.
func DBConn(cfg string) (*DB, error) {
	db, err := sql.Open(consts.Driver, cfg)
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
	if _, err := db.DB.Exec(consts.InsertQuery, sub.ServiceName, sub.Price, sub.UserID, sub.StartDate); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			err := errors.New(consts.AlreadyExist)
			return errors.WithStack(err)
		}
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetSubscription(sub *structs.Subscription) (*structs.Subscription, error) {
	row := db.DB.QueryRow(consts.SelectQuery, sub.UserID, sub.ServiceName)
	var dbRow structs.Subscription
	if err := row.Scan(&dbRow.ServiceName, &dbRow.Price, dbRow.UserID, dbRow.StartDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WithStack(err)
		}
		return nil, errors.WithStack(err)
	}
	return &dbRow, nil
}

func (db *DB) UpdateSubscription(update *structs.Subscription) (*sql.Result, error) {
	row, err := db.Exec(consts.UpdateQuery, update.UserID, update.ServiceName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &row, nil
}
