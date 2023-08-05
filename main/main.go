package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Arpit-Mohapatra/urlshort"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://www.linkedin.com/in/arpit-mohapatra-708520248/",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	
	yamlFile := flag.String("yaml", "data.yaml", "yaml file in the form of path and url")
	flag.Parse()
	yamlData, err := os.ReadFile(*yamlFile)
	if err != nil {
		fmt.Printf("Could not read %v", *yamlFile)
		return
	}
	yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
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

