package api

import "net/http"

type Book struct {
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type Books map[string]Book

func addBook(request *http.Request, writer http.ResponseWriter) {

}

func listBooks(writer http.ResponseWriter) {

}

func deleteBook(isbn string, writer http.ResponseWriter) {

}

func updateBook(isbn string, writer http.ResponseWriter) {

}

func getBook(isbn string, writer http.ResponseWriter) {

}
