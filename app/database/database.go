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

// Функция реализует операцию "get" получая строку из базы данных.
func (db *DB) GetSubscription(sub *structs.Subscription) (*structs.Subscription, error) {
	row := db.DB.QueryRow(consts.SelectQuery, sub.ServiceName, sub.UserID)
	var result structs.Subscription
	if err := row.Scan(&result.ServiceName, &result.Price, result.UserID, result.StartDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WithStack(err)
		}
		return nil, errors.WithStack(err)
	}
	return &result, nil
}

// Функция реализует операцию "update" обновляя существующую в базе данных строку.
// // Проверяет наличие строки и в случае ее присутствия обновляет поле "price".
func (db *DB) UpdateSubscription(sub *structs.Subscription) (*structs.Subscription, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	row := tx.QueryRow(consts.UpdateQuery, sub.ServiceName, sub.UserID)
	var result structs.Subscription
	if err = row.Scan(&result.ServiceName, &result.Price, &result.UserID, &result.StartDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WithStack(err)
		}
		return nil, errors.WithStack(err)
	}

	return &result, nil
}

func (db *DB) DeleteSubscription(sub *structs.Subscription) error {
	result, err := db.Exec(consts.DeleteQuery, sub.ServiceName, sub.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errors.WithStack(err)
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
