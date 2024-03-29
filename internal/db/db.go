package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(connectionString string) {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
