package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

type DB struct {
	*sql.DB
}

func NewDB(cfg string) (*DB, error) {
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
}
