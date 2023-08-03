package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8080")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from port 8080")
}