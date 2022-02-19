package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(dbName, login, password, url, port string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", login, password, url, port, dbName))
	if err != nil {
		return nil, err
	}

	fmt.Println("t", db.Ping())
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
