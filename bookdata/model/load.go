package model

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type Book struct {
	BookID         string  `json:"book_id"`
	Title          string  `json:"title"`
	Authors        string  `json:"authors"`
	AverageRatings float64 `json:"average_rating"`
	ISBN           string  `json:"isbn"`
	ISBN13         string  `json:"isbn_13"`
	LanguageCode   string  `json:"language_code"`
	NumPages       int     `json:"num_panges"`
	Ratings        int     `json:"ratings"`
	Reviews        int     `json:"reviews"`
}

// loadData is use to read csv and returns a pointer to the slice of Books
func loadData(r io.Reader) *[]*Book {
	reader := csv.NewReader(r)
	var res []*Book
	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of file.")
			break
		} else if err != nil {
			log.Println("Error occured while reading the file : ", err)
			break
		}
		averageRating, err := strconv.ParseFloat(row[3], 64)
		numPages, err := strconv.Atoi(row[7])
		ratings, err := strconv.Atoi(row[8])
		reviews, err := strconv.Atoi(row[9])
		if err != nil {
			log.Println("Error in parsing some numbers : ", err)
		}
		book := &Book{
			BookID:         row[0],
			Title:          row[1],
			Authors:        row[2],
			AverageRatings: averageRating,
			ISBN:           row[4],
			ISBN13:         row[5],
			LanguageCode:   row[6],
			NumPages:       numPages,
			Ratings:        ratings,
			Reviews:        reviews,
		}
		res = append(res, book)
	}
	return &res
}
