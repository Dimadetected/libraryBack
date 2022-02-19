package repositories

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repositories struct {
	db *sqlx.DB
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		db: db,
	}
}
