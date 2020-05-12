package main

import (
	"log"
	"net/http"

	"github.com/farzamalam/gopher-exercises/bookdata/handler"
	"github.com/farzamalam/gopher-exercises/bookdata/util"
	"github.com/gorilla/mux"
)

// TO DO
// 0. Read and understand models, operations, urls and handler funcs  --> Done.
// 1. Make model and data object to hold data.	--> Done.
// 2. Implement operations on slice of books`  --> Done.
// 3. Write urls and query paramters.
// 4. Implement all the handler funcs, and exopose operations
// 5. Test on localhost
// 5. Deploy on Nginx Server.

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", home)
	api.HandleFunc("/books/isbn/{isbn}", handler.SearchByISBN).Methods(http.MethodGet)
	log.Println("Starting server at : 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func home(w http.ResponseWriter, r *http.Request) {
	data := util.Message(true, "Welcome home!")
	util.Respond(w, http.StatusOK, data)
}
