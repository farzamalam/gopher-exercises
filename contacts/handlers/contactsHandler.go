package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/farzamalam/gopher-exercises/contacts/utils"
	"github.com/gorilla/mux"

	"github.com/farzamalam/gopher-exercises/contacts/models"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	// TODO : Take the user id from the context.
	contact := models.Contact{}
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, utils.Message(false, "Error in the message body."))
		return
	}
	resp := contact.Create()
	utils.Respond(w, http.StatusCreated, resp)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	// TODO : Take the user id from the context
	// Get the userID from the url.
	// Call the GetContacts() of models and check for errors.
	// Handle errors and response accordingly.
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message(false, "Error in the userID Param"))
		return
	}
	contacts, err := models.GetContacts(userID)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, utils.Message(false, "Error while Getting the contacts"))
		return
	}
	if contacts == nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message(false, "No contacts are found."))
		return
	}
	resp := utils.Message(true, "Success")
	resp["data"] = contacts
	utils.Respond(w, http.StatusOK, resp)
}
