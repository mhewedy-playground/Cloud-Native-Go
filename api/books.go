package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var booksDb = make(map[string]*Book)

func addBook(r *http.Request, w http.ResponseWriter) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	book, err := jsonToBook(bytes)
	if err != nil {
		handleError(err, w)
		return
	}

	booksDb[book.ISBN] = book
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", "/api/books/"+book.ISBN)
}

func listBooks(w http.ResponseWriter) {

	var books []*Book

	for _, book := range booksDb {
		books = append(books, book)
	}

	bytes, err := booksToJson(books)
	if err != nil {
		handleError(err, w)
		return
	}
	_, _ = w.Write(bytes)
}

func deleteBook(isbn string, w http.ResponseWriter) {

}

func updateBook(isbn string, w http.ResponseWriter) {

}

func getBook(isbn string, w http.ResponseWriter) {

}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(err.Error()))
	log.Println(err.Error())
}
