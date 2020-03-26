package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Story map[string]Chapter

func JsonDecode(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	err := d.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultTemplate))
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v\n", err)
			http.Error(w, "Error while Executing the template", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Path not found", http.StatusNotFound)
}

var defaultTemplate = `<!DOCTYPE html>
			<html>
				<head>
					<meta charset="utf-8">
					<tilte>Choose Your own adventure</tilte>
				</head>

				<body>
					<h1>{{.Title}} </h1>
					{{range .Paragraphs}}
					<p>{{.}} </p>
					{{end}}

					<ul>
						{{range .Options}}
						<li><a href="/{{.Chapter}}">{{.Text}}</a> </li>
						{{end}}
					</ul>
				</body>
			</html>`
