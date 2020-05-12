package handler

import (
	"net/http"

	"github.com/farzamalam/gopher-exercises/bookdata/model"
	"github.com/farzamalam/gopher-exercises/bookdata/util"
	"github.com/gorilla/mux"
)

var books model.BookStore

func init() {
	books = &model.Books{}
	books.Initialize()

}
func SearchByISBN(w http.ResponseWriter, r *http.Request) {
	isbn, ok := mux.Vars(r)["isbn"]
	if !ok {
		data := util.Message(false, "ISBN not found in the url.")
		util.Respond(w, http.StatusNotFound, data)
		return
	}
	data := books.SearchISBN(isbn)
	if data != nil {
		resp := util.Message(true, "Success")
		resp["data"] = data
		util.Respond(w, http.StatusOK, resp)
		return
	}
	resp := util.Message(false, "Not found")
	util.Respond(w, http.StatusNotFound, resp)
}
