package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/farzamalam/gopher-exercises/go-link-parser/links"
)

func main() {
	// Created a flag url with default "https://www.gophercises.com"
	urlFlag := flag.String("url", "https://gophercises.com", "the url to make sitemap for.")
	flag.Parse()
	fmt.Println(*urlFlag)

	hrefs := get(*urlFlag)

	for i, href := range hrefs {
		fmt.Printf("%d :  %s\n", i, href)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(base, href(resp.Body, base))
}

func href(r io.Reader, base string) []string {
	var res []string
	links, _ := links.Parse(r)
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			res = append(res, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			res = append(res, link.Href)
		}
	}
	return res
}

func filter(base string, hrefs []string) []string {
	var filterHref []string
	for _, href := range hrefs {
		if strings.HasPrefix(href, base) {
			filterHref = append(filterHref, href)
		}
	}
	return filterHref
}
