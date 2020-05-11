package model

import (
	"fmt"
	"log"
	"os"
)

// Initialize opens the csv and calls the loadData to read the csv and gets a book that contains
// sile of pointer to book
func Initialize() {
	fileName := "books.csv"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error while opening the file %s\n", fileName)
	}
	books := loadData(file)
	for _, book := range *books {
		fmt.Println(*book)
	}
}
