package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/farzamalam/gopher-exercises/contacts/models"
	"github.com/farzamalam/gopher-exercises/contacts/utils"
	"github.com/gorilla/mux"
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

func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Endpoint : /api/v1/user/{accountsID}
	params := mux.Vars(r)
	accountID, err := strconv.Atoi(params["accountsID"])
	if err != nil {
		utils.Respond(w, http.StatusBadRequest, utils.Message(false, "Invalid accountsID in the url"))
		return
	}
	account := models.GetUser(accountID)
	if account == nil {
		utils.Respond(w, http.StatusInternalServerError, utils.Message(false, "Error while Getting the user data."))
		return
	}
	resp := utils.Message(true, "Success")
	resp["data"] = account
	utils.Respond(w, http.StatusOK, resp)
}
