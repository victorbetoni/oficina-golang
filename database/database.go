package database

import (
	"os"

	"github.com/jmoiron/sqlx"
)

var db sqlx.DB

func Connect() (err error) {

	var database *sqlx.DB
	if database, err = sqlx.Open("postgres", os.Getenv("DB_URI")); err != nil {
		return err
	}
	db = *database
	return
}

func GrabDB() *sqlx.DB {
	if err := db.Ping(); err != nil {
		Connect()
	}
	return &db
}
