package models

// Contact type is used to store single contacts in contacts table.
type Contact struct {
	ContactsID uint   `json:"contacts_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	UserID     uint   `json:"user_id"`
	CreatedID  string `json:"created_at"`
}
