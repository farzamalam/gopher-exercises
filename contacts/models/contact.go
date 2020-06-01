package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/farzamalam/gopher-exercises/contacts/utils"
)

type Contact struct {
	ContactsID int    `json:"contacts_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	UserID     int    `json:"user_id"`
	CreatedAt  string `json:"created_at"`
}

func GetContact(id int) *Contact {
	contact := Contact{}
	err := GetDB().QueryRow("Select contacts_id, name, phone, user_id, created_at from contacts where user_id = ? ", id).Scan(&contact.ContactsID, &contact.Name, &contact.Phone, &contact.UserID, &contact.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No Record found for id ", id)
			return nil
		}
		log.Fatal("Error whlie Getting Contact : ", err)
	}
	return &contact
}

func GetContacts(id int) []*Contact, error {
	var contacts []*Contact
	res, err := GetDB().Query("Select contacts_id, name, phone, user_id, created_at from contacts where user_id = ? ", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No Record found for id ", id)
			return nil,nil
		}
		log.Println("Error while Getting Contacts : ", err)
		return nil, err
	}
	for res.Next() {
		contact := Contact{}
		err = res.Scan(&contact.ContactsID, &contact.Name, &contact.Phone, &contact.UserID, &contact.CreatedAt)
		if err != nil {
			log.Println("Error while Calling the service : ", err)
			nil, err
		}
		contacts = append(contacts, &contact)
	}
	return contacts, nil
}

func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}
	sql := `
		INSERT INTO contacts (name, phone , user_id)
		VALUES(?,?,?) ;
	`
	_, err := GetDB().Query(sql, contact.Name, contact.Phone, contact.UserID)
	if err != nil {
		log.Println("Eror while Inserting Contact : ", err)
		return utils.Message(false, "Error while Inserting Contact")

	}
	err = GetDB().QueryRow("Select * from contacts where contacts_id = (SELECT MAX(contacts_id) from contacts)").Scan(&contact.ContactsID, &contact.Name, &contact.Phone, &contact.UserID, &contact.CreatedAt)
	if err != nil {
		log.Println("Error while getting the contact details after insert", err)
		return utils.Message(false, "Error  while getting the contact details after insert")
	}
	resp := utils.Message(true, "Success")
	fmt.Println("contact", contact)
	resp["data"] = contact
	return resp
}

func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return utils.Message(false, "Invalid Contact Name"), false
	}
	if contact.Phone == "" {
		return utils.Message(false, "Invalid Phone"), false
	}
	if contact.UserID <= 0 {
		return utils.Message(false, "Invalid UserID"), false
	}
	// All required paramters are met.
	return utils.Message(true, "Sucess"), true
}
