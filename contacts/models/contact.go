package models

type Contact struct {
	ContactsID int    `json:"contacts_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	UserID     int    `json:"user_id"`
	CreatedAt  string `json:"created_at"`
}
