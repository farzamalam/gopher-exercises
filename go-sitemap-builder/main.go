package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Created a flag url with default "https://www.gophercises.com"
	urlFlag := flag.String("url", "https://gophercises.com", "the url to make sitemap for.")
	flag.Parse()
	fmt.Println(*urlFlag)
	// Call a GET request on the url and copy the resp body on os.Stdout and close the resp.
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
