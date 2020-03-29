package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Password@123"
	dbName   = "gophercises_phone"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	pssqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s sslmode = disable ", host, port, user, password)
	// db, err := sql.Open("postgres", pssqlInfo)
	// must(err)
	// must(resetDb(db, dbName))
	// defer db.Close()

	pssqlInfo = fmt.Sprintf("%s dbname = %s", pssqlInfo, dbName)
	db, err := sql.Open("postgres", pssqlInfo)
	must(err)
	defer db.Close()
	must(createPhoneNumberTable(db))

	id, err := insertPhoneNumber(db, "1234567890")
	must(err)
	fmt.Println("id = ", id)

}

func insertPhoneNumber(db *sql.DB, phone string) (int, error) {
	statement := `
	INSERT INTO PHONE_NUMBER(value) VALUES($1) RETURNING ID
	`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	return id, err
}
func createPhoneNumberTable(db *sql.DB) error {
	statement := `
	CREATE TABLE IF NOT EXISTS PHONE_NUMBER (
	ID SERIAL,
	VALUE VARCHAR(255)
	)
	`
	_, err := db.Exec(statement)
	return err
}

func resetDb(db *sql.DB, name string) error {

	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		panic(err)
	}
	return createDb(db, name)
}

func createDb(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	return err
}
