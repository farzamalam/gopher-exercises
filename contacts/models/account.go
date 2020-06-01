package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/farzamalam/gopher-exercises/contacts/utils"
)

type Token struct {
	UserID int
	jwt.StandardClaims
}

type Account struct {
	AccountsID int    `json:"accounts_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Token      string `json:"token";sql:"-"`
	CreatedAt  string `json:"created_at"`
}

// Validate incoming user details.
func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}
	if len(account.Password) < 6 {
		return utils.Message(false, "Valid Password is required"), false
	}
	// Email must be unique
	row, err := GetDB().Query("Select accounts_id from accounts where email = ?", account.Email)
	fmt.Println("account.Email : ", account.Email)
	var id int
	row.Scan(&id)
	fmt.Println("row : ", id)
	if err == nil {
		return utils.Message(false, "Email address is already present."), false
	}
	if err != sql.ErrNoRows {
		return utils.Message(false, "Connection Error. Please retry"), false
	}

	return utils.Message(false, "Requirement Passed."), true
}

// Create new account.
func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	sql := `
		Insert into accounts (name, email, password)
		values(?,?,?);
	`
	_, err := GetDB().Query(sql, account.Name, account.Email, account.Password)
	if err != nil {
		return utils.Message(false, "Error while Inserting the account")
	}
	sql = "select accounts_id, name, email,password, created_at  from accounts where accounts_id= (select max(accounts_id) from accounts)"
	err = GetDB().QueryRow(sql).Scan(&account.AccountsID, &account.Name, &account.Email, &account.Password, &account.CreatedAt)
	if err != nil {
		log.Println("Error while getting the Account details after the insert : ", err)
		return utils.Message(false, "Error while getting the Account details after insert.")
	}

	// Create a new JTW token for the newly registered account.
	tk := &Token{UserID: account.AccountsID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = ""
	resp := utils.Message(true, "Account has been created.")
	resp["data"] = account
	return resp
}
