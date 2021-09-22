package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type BasicAuthResponse struct {
	Authenticated bool   `json:"authenticated"`
	User          string `json:"user"`
}

type CreateAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAuthResponse struct {
	Created bool   `json:"created"`
	User    string `json:"user"`
}

var db *sql.DB

func init() {
	_ = godotenv.Load()
	name := os.Getenv("db_name")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	user := os.Getenv("db_user")
	pass := os.Getenv("db_pass")

	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name))
	if err != nil {
		log.Fatalf("error in connecting to database: %v", err)
	}
}

func main() {

	defer db.Close()
	http.HandleFunc("/api/v1/verify", verify)
	http.HandleFunc("/api/v1/create", create)

	port := "8080"
	log.Printf("Starting server at : %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func verify(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		log.Printf("Empty credentials\n")
		fmt.Fprintf(w, "Empty credentials\n")
		return
	}
	log.Printf("Username: %s\n", username)
	log.Printf("Password: %s\n", password)

	v := verifyInDB(username, password)
	if v {
		out := BasicAuthResponse{
			Authenticated: v,
			User:          username,
		}
		resp, _ := json.Marshal(out)
		fmt.Fprint(w, string(resp))
		return
	}

}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Invalid method\n")
		return
	}
	var b CreateAuth

	err := json.NewDecoder(r.Body).Decode(&b)
	defer r.Body.Close()

	if err != nil {
		fmt.Fprintf(w, "Invalid body: %s\n", err)
		return
	}
	log.Printf("Username: %s\n", b.Username)
	log.Printf("Password: %s\n", b.Password)
	err = createInDB(b.Username, b.Password)
	if err != nil {
		fmt.Fprintf(w, "Internal server Error: %s", err)
		return
	}
	data := CreateAuthResponse{
		Created: true,
		User:    b.Username,
	}
	resp, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s\n", resp)
}

func verifyInDB(username, password string) bool {
	res := db.QueryRow("select password from auth_table where username = ?", username)
	var pass string
	err := res.Scan(&pass)
	if err != nil {
		log.Printf("error while verifying: %v\n", err)
		return false
	}
	if pass != password {
		return false
	}
	return true
}
func createInDB(username, password string) error {
	sql := fmt.Sprintf("insert into auth_table(username, password) values('%s', '%s')", username, password)
	_, err := db.Exec(sql)
	if err != nil {
		log.Printf("error while create user: %v\n", err)
		return err
	}
	return nil
}
