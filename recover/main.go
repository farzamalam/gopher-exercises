package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/panic", panicDemo)
	mux.HandleFunc("/panic-after", panicAfterDemo)
	log.Fatal(http.ListenAndServe(":3000", recoverMW(mux)))
}

func recoverMW(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				http.Error(w, "Something went wrong :(", http.StatusInternalServerError)
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello !</h1>")
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanic()
}

func funcThatPanic() {
	panic("oh no!")
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello !</h1>")
	funcThatPanic()
}
