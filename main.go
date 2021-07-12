package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// lets rocket!

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	log.Println("Listing for requests as http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
