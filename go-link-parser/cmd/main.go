package main

import (
	"fmt"
	"strings"

	"github.com/farzamalam/gopher-exercises/go-link-parser/links"
)

var exampleHtml = `<html>
					<body>
						<h1>Hello!</h1>
						<a href="/other-page">A link to another page</a>
					</body>
				</html>`

func main() {
	r := strings.NewReader(exampleHtml)

	links, err := links.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
