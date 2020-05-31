package models

import (
	"database/sql"
	"log"
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

func GetContacts(id int) []*Contact {
	var contacts []*Contact
	res, err := GetDB().Query("Select contacts_id, name, phone, user_id, created_at from contacts where user_id = ? ", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No Record found for id ", id)
			return nil
		}
		log.Fatal("Error while Getting Contacts : ", err)
	}
	for res.Next() {
		contact := Contact{}
		err = res.Scan(&contact.ContactsID, &contact.Name, &contact.Phone, &contact.UserID, &contact.CreatedAt)
		if err != nil {
			log.Fatal("Error while Calling the service : ", err)
		}
		contacts = append(contacts, &contact)
	}
	return contacts
}
