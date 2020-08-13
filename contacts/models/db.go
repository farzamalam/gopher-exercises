package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

var db *sql.DB

// init is used to initialize the .env file and open a connection to db.
func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error while Loading the environment variables : ", e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, dbHost, dbName)
	log.Println("dbURI : ", dbURI)
	var err error
	db, err = sql.Open(dbType, dbURI)
	if err != nil {
		log.Fatal("Error while connecting to Database : ", err)
	}

}

// GetDB returns the DB type that is used to close the connection from main.
func GetDB() *sql.DB {
	return db
}
