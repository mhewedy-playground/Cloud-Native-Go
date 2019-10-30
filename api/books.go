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

func addBook(w http.ResponseWriter, r *http.Request) {
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
	if _, found := booksDb[isbn]; !found {
		notFound(isbn, w)
		return
	}

	delete(booksDb, isbn)
}

func updateBook(isbn string, w http.ResponseWriter, r *http.Request) {

	if _, found := booksDb[isbn]; !found {
		notFound(isbn, w)
		return
	}

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

	booksDb[isbn] = book
}

func getBook(isbn string, w http.ResponseWriter) {
	book, found := booksDb[isbn]

	if !found {
		notFound(isbn, w)
		return
	}

	bytes, err := bookToJson(book)
	if err != nil {
		handleError(err, w)
		return
	}
	_, _ = w.Write(bytes)
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(err.Error()))
	log.Println(err.Error())
}

func notFound(isbn string, w http.ResponseWriter) {
	strErr := "isbn " + isbn + " not found"
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(strErr))
	log.Println(strErr)
}
