package urlshort

import (
	"net/http"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusOK)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}