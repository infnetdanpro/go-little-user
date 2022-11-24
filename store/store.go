package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	dsn := "postgresql://postgres:postgres@localhost:5432/gotest"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	return db
}
