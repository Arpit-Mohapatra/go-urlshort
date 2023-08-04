package main

import (
	"fmt"
	"net/http"

	"github.com/Arpit-Mohapatra/urlshort"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://www.linkedin.com/in/arpit-mohapatra-708520248/",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	yaml := `
- path: /urlshort
  url: https://github.com/Arpit-Mohapatra/go-urlshort
- path: /urlshort-final
  url: https://github.com/Arpit-Mohapatra/
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting sever at :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from port 8080")
}