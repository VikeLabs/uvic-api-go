package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
)

func init() {
	var err error
	db, err = sqlx.Connect("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
}

func New() *sqlx.DB {
	return db
}
