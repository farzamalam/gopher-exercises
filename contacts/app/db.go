package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	e := godotenv.Load() // load .env file
	if e != nil {
		log.Fatalln(e)
	}

	pssqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s sslmode = disable ", host, port, user, password)
	log.Println("pssqlInfo : ", pssqlInfo)
	db, err := sql.Open()
	defer db.Close()
}
