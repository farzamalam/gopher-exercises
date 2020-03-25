package handler

import "net/http"

func MapHandler(pathUrls map[string]string, fallback http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusOK)
			return
		}
	}
	fallback.ServeHTTP(w, r)
}
