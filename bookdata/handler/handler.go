package handler

import (
	"encoding/json"
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

func SearchByBookName(w http.ResponseWriter, r *http.Request) {
	bookName, ok := mux.Vars(r)["book"]
	if !ok {
		data := util.Message(false, "Invalid book name")
		util.Respond(w, http.StatusNotFound, data)
		return
	}
	ratingOver, ratingBelow, err := getRatingParams(r)
	limit, err := getLimitParam(r)
	skip, err := getSkipParam(r)
	if err != nil {
		data := util.Message(false, "Invalid query params")
		util.Respond(w, http.StatusNotFound, data)
		return
	}
	data := books.SearchBook(bookName, ratingOver, ratingBelow, limit, skip)
	resp := util.Message(true, "Success")
	resp["data"] = data
	util.Respond(w, http.StatusOK, resp)

}

func SearchByAuthor(w http.ResponseWriter, r *http.Request) {
	author, ok := mux.Vars(r)["author"]
	if !ok {
		data := util.Message(false, "Invalid author name")
		util.Respond(w, http.StatusNotAcceptable, data)
		return
	}
	ratingOver, ratingBelow, err := getRatingParams(r)
	limit, err := getLimitParam(r)
	skip, err := getSkipParam(r)
	if err != nil {
		data := util.Message(false, "Invalid query param")
		util.Respond(w, http.StatusNotAcceptable, data)
		return
	}
	data := books.SearchAuthor(author, ratingOver, ratingBelow, limit, skip)
	resp := util.Message(true, "Success")
	resp["data"] = data
	util.Respond(w, http.StatusOK, resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book *model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	defer r.Body.Close()
	if err != nil {
		data := util.Message(false, "Invalid body")
		util.Respond(w, http.StatusNotAcceptable, data)
		return
	}
	ok := books.CreateBook(book)
	resp := util.Message(ok, "New Book is created.")
	resp["data"] = book
	util.Respond(w, http.StatusCreated, resp)
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
