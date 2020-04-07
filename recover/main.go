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
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
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
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	if err != nil {
		line = -1
	}
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
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, b.String())
	style := styles.Get("github")
	if style == nil {
		style = styles.Fallback
	}
	formatter := html.New(html.TabWidth(2), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)
	// _ = quick.Highlight(w, b.String(), "go", "html", "github")
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

func makeLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	for li, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		file := ""
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
		var lineStr strings.Builder
		for i := len(file) + 2; i < len(line); i++ {
			if line[i] < '0' || line[i] > '9' {
				break
			}
			lineStr.WriteByte(line[i])
		}
		v := url.Values{}
		v.Set("path", file)
		v.Set("line", lineStr.String())
		lines[li] = "\t<a href=\"/debug/?" + v.Encode() + "\">" + file + ":" + lineStr.String() + "</a>" + line[len(file)+2+len(lineStr.String()):]
	}
	return strings.Join(lines, "\n")
}
