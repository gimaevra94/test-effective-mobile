// Пакет предоставляет модуль для взаимодействия с базой данных.
package database

import (
	"database/sql"
	"time"

	"github.com/gimaevra94/test-effective-mobile/app/consts"
	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type DB struct {
	*sql.DB
}

// Функция открывает соединение с базой данных.
// Пингует его чтобы проверить готовность принимать запросы.
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
// Сканирует полученные поля в структуру и отдает вызывающей функции.
func (db *DB) GetSubscription(sub *structs.Subscription) (structs.Subscription, error) {
	row := db.DB.QueryRow(consts.SelectQuery, sub.ServiceName, sub.UserID)
	var result structs.Subscription
	if err := row.Scan(&result.ServiceName, &result.Price, &result.UserID, &result.StartDate); err != nil {
		if err == sql.ErrNoRows {
			return structs.Subscription{}, errors.WithStack(err)
		}
		return structs.Subscription{}, errors.WithStack(err)
	}
	return result, nil
}

// Функция реализует операцию "update" обновляя существующую в базе данных строку.
// В одном запросе происходит операция обновления и возврата строки.
// По этому сначала открывается транзакция и через нее выполняется запрос.
// Поля возвращенной строки сканируются в структуру которая возвращается вызывающей функции.
func (db *DB) UpdateSubscription(sub *structs.Subscription) (structs.Subscription, error) {
	tx, err := db.Begin()
	if err != nil {
		return structs.Subscription{}, errors.WithStack(err)
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	row := tx.QueryRow(consts.UpdateQuery, sub.Price, sub.ServiceName, sub.UserID)
	var result structs.Subscription
	if err = row.Scan(&result.ServiceName, &result.Price, &result.UserID, &result.StartDate); err != nil {
		if err == sql.ErrNoRows {
			return structs.Subscription{}, errors.WithStack(err)
		}
		return structs.Subscription{}, errors.WithStack(err)
	}

	if err = tx.Commit(); err != nil {
		return structs.Subscription{}, errors.WithStack(err)
	}

	return result, nil
}

// Функция реализует операцию 'delete' удаляя строку по ключу "service_name + user_id".
// Так же запрашивается RowsAffected из result чтобы понять происзошло ли удаление.
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

func (db *DB) GetPeriodTotalPrice(serviceName, userID string, fromDate time.Time) (int, error) {
	row := db.QueryRow(consts.GetTotalPriceSelectQuery, serviceName, userID, fromDate)
	var totalPrice int
	if err := row.Scan(&totalPrice); err != nil {
		if err == sql.ErrNoRows {
			err := errors.New(consts.NotExist)
			return 0, errors.WithStack(err)
		}
		return 0, errors.WithStack(err)
	}
	return totalPrice, nil
}
