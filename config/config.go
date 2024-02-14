package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var TableName string

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "root"
	dbName   = "enigma_laundry"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}