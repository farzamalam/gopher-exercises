package handler

import (
	"net/http"
	"strconv"

	"github.com/farzamalam/gopher-exercises/bookdata/model"
	"github.com/farzamalam/gopher-exercises/bookdata/util"
	"github.com/gorilla/mux"
)

var books model.BookStore

func init() {
	books = &model.Books{}
	books.Initialize()

}

func Home(w http.ResponseWriter, r *http.Request) {
	data := util.Message(true, "Welcome home!")
	util.Respond(w, http.StatusOK, data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	isbn, ok := mux.Vars(r)["isbn"]
	if !ok {
		data := util.Message(false, "ISBN is not found in the url.")
		util.Respond(w, http.StatusNotFound, data)
		return
	}
	ok = books.DeleteBook(isbn)
	if !ok {
		data := util.Message(false, "No record is found.")
		util.Respond(w, http.StatusNotFound, data)
		return
	}
	data := util.Message(true, "Record has been deleted.")
	util.Respond(w, http.StatusAccepted, data)
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

// getRatingParams is used to get the rating params from the optional query parameter.
func getRatingParams(r *http.Request) (float64, float64, error) {
	ratingOver := 0.0
	ratingBelow := 5.0
	query := r.URL.Query()
	ro := query.Get("ratingOver")
	if ro != "" {
		val, err := strconv.ParseFloat(ro, 64)
		if err != nil {
			return ratingOver, ratingBelow, err
		}
		ratingOver = val
	}
	rb := query.Get("ratingBelow")
	if rb != "" {
		val, err := strconv.ParseFloat(rb, 64)
		if err != nil {
			return ratingOver, ratingBelow, err
		}
		ratingBelow = val
	}
	return ratingOver, ratingBelow, nil
}

func getLimitParam(r *http.Request) (int, error) {
	limit := 0
	query := r.URL.Query()
	l := query.Get("limit")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return limit, err
		}
		limit = val
	}
	return limit, nil
}

func getSkipParam(r *http.Request) (int, error) {
	skip := 0
	query := r.URL.Query()
	s := query.Get("skip")
	if s != "" {
		val, err := strconv.Atoi(s)
		if err != nil {
			return skip, err
		}
		skip = val
	}
	return skip, nil
}
