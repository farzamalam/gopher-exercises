package model

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type BookStore interface {
	Initialize()
	PrintBooks()
	Len() int
	SearchAuthor(author string, ratingOver, ratingBelow float64, limit, skip int) *[]*Book
	SearchBook(bookName string, ratingOver, ratingBelow float64, limit, skip int) *[]*Book
	SearchISBN(isbn string) *Book
	CreateBook(book *Book) bool
	DeleteBook(isbn string) bool
	UpdateBook(isbn string, book *Book) bool
}

type Books struct {
	Store *[]*Book
}

// Initialize opens the csv and calls the loadData to read the csv and gets a book that contains
// sile of pointer to book
func (b *Books) Initialize() {
	fileName := "books.csv"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error while opening the file %s\n", fileName)
	}
	b.Store = loadData(file)
}
func (b *Books) SearchAuthor(author string, ratingOver, ratingBelow float64, limit, skip int) *[]*Book {
	res := Filter(b.Store, func(bk *Book) bool {
		return strings.Contains(strings.ToLower(bk.Authors), strings.ToLower(author)) && bk.AverageRatings >= ratingOver && bk.AverageRatings <= ratingBelow
	})
	if limit == 0 || limit > len(*res) {
		limit = len(*res)
	}
	data := (*res)[skip:limit]
	return &data
}
func (b *Books) SearchBook(bookName string, ratingOver, ratingBelow float64, limit, skip int) *[]*Book {
	res := Filter(b.Store, func(bk *Book) bool {
		return strings.Contains(strings.ToLower(bk.Title), strings.ToLower(bookName)) && bk.AverageRatings >= ratingOver && bk.AverageRatings <= ratingBelow
	})
	if limit == 0 || limit > len(*res) {
		limit = len(*res)
	}
	data := (*res)[skip:limit]
	return &data
}

func (b *Books) UpdateBook(isbn string, book *Book) bool {
	for _, bk := range *b.Store {
		if bk.ISBN == isbn {
			bk = book
			return true
		}
	}
	return false
}

func (b *Books) CreateBook(book *Book) bool {
	*b.Store = append(*b.Store, book)
	return true
}

func (b *Books) DeleteBook(isbn string) bool {
	idx := -1
	for i, bk := range *b.Store {
		if bk.ISBN == isbn {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false
	}
	res := (*b.Store)[:idx]
	res = append(res, (*b.Store)[idx+1:]...)
	*b.Store = res
	return true
}

func (b *Books) PrintBooks() {
	for _, book := range *b.Store {
		fmt.Println(*book)
	}
}

func (b *Books) Len() int {
	return len(*b.Store)
}

func (b *Books) SearchISBN(isbn string) *Book {
	res := Filter(b.Store, func(book *Book) bool {
		return strings.ToLower(book.ISBN) == strings.ToLower(isbn)
	})
	if len(*res) > 0 {
		return (*res)[0]
	}
	return nil
}

func Filter(books *[]*Book, f func(*Book) bool) *[]*Book {
	var res []*Book
	for _, book := range *books {
		if f(book) {
			res = append(res, book)
		}
	}
	return &res
}
