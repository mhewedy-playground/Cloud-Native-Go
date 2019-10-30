package api

import (
	"net/http"
	"strings"
)

func HandleBooks(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		listBooks(w)
	case http.MethodPost:
		addBook(r, w)
	default:
		unsupportedHttpMethod(w)
	}
}

func HandleBook(w http.ResponseWriter, r *http.Request) {
	isbn := strings.TrimPrefix(r.URL.Path, "/api/book/")

	switch r.Method {
	case http.MethodGet:
		getBook(isbn, w)
	case http.MethodPut:
		updateBook(isbn, w)
	case http.MethodDelete:
		deleteBook(isbn, w)
	default:
		unsupportedHttpMethod(w)
	}
}

func unsupportedHttpMethod(resp http.ResponseWriter) {
	_, _ = resp.Write([]byte("unsupported http method"))
	resp.WriteHeader(http.StatusBadRequest)
}
