package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	e := godotenv.Load() // load .env file
	if e != nil {
		log.Fatalln(e)
	}

	pssqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s sslmode = disable ", host, port, user, password)
	log.Println("pssqlInfo : ", pssqlInfo)
	db, err := sql.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
