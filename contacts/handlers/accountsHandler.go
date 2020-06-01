package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farzamalam/gopher-exercises/contacts/models"
	"github.com/farzamalam/gopher-exercises/contacts/utils"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	defer r.Body.Close()
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message(false, "Invalid Request Body."))
		return
	}
	resp := account.Create()
	utils.Respond(w, http.StatusCreated, resp)
}
