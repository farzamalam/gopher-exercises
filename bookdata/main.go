package main

import (
	"fmt"

	"github.com/farzamalam/gopher-exercises/bookdata/model"
)

// TO DO
// 0. Read and understand models, operations, urls and handler funcs  --> Done.
// 1. Make model and data object to hold data.  --> Done.
// 2. Implement operations on slice of books
// 3. Write urls and query paramters.
// 4. Implement all the handler funcs, and exopose operations
// 5. Deploy on Nginx Server.

var books model.BookStore

func init() {
	books = &model.Books{}
	books.Initialize()

}

func main() {
	books.PrintBooks()
	fmt.Println("length : ", books.Len())
}
