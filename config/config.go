package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	TableName string
	host     string
	port     string
	user     string
	password string
	dbName   string
	apiPort  string
)

func readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	apiPort = os.Getenv("API_PORT")

	return nil
}

func ConnectDB() *sql.DB {
	err := readConfig()
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

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

func TestConnectDB() {
	err := readConfig()
	if err != nil {
		panic(err)
	}
	
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	fmt.Println(psqlInfo)
}

func ApiPort() string {
	err := readConfig()

	if err != nil {
		panic(err)
	}
	
	return apiPort
}