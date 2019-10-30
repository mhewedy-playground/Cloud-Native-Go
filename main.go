package main

import (
	"Cloud-Native-Go/api"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/api/books", api.HandleBooks)
	http.HandleFunc("/api/book/", api.HandleBook)

	log.Fatal(http.ListenAndServe(port(), nil))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
