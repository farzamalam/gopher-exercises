package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strings"

	"github.com/alecthomas/chroma/quick"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/debug/", sourceCodeHandler)
	mux.HandleFunc("/panic", panicDemo)
	mux.HandleFunc("/panic-after", panicAfterDemo)
	log.Fatal(http.ListenAndServe(":3000", recoverMW(mux, true)))
}

func recoverMW(app http.Handler, dev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				stack := debug.Stack()
				log.Println(string(stack))
				if !dev {
					http.Error(w, "Something went wrong :(", http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "<h1>panic : %v </h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()
		// nw := &responseWriter{
		// 	ResponseWriter: w,
		// }
		app.ServeHTTP(w, r)
		//nw.flush()
	}
}

// type ResponseWriter interface {
// 	Header() Header
// 	Write([]byte) (int, error)
// 	WriteHeader(statusCode int)
// }
func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//err = quick.Highlight(os.Stdout, "package main", "go", "html", "monokai")
	err = quick.Highlight(w, b.String(), "go", "html", "github")
}

func makeLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	for li, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		var file string
		count := 0
		for i, ch := range line {
			if ch == ':' {
				count++
				if count == 2 {
					file = line[1:i]
					count = 0
					break
				}

			}
		}
		v := url.Values{}
		v.Set("path", file)
		lines[li] = "\t <a href = \"/debug/?" + v.Encode() + "\">" + file + "</a>" + line[len(file)+1:]
	}
	return strings.Join(lines, "\n")
}

type responseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.writes = append(rw.writes, b)
	return len(b), nil
}
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
}

func (rw *responseWriter) flush() error {
	if rw.status != 0 {
		rw.ResponseWriter.WriteHeader(rw.status)
	}
	for _, write := range rw.writes {
		_, err := rw.ResponseWriter.Write(write)
		if err != nil {
			return err
		}
	}
	return nil
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
