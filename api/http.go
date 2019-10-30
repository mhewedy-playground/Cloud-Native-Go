package api

import (
	"net/http"
	"strings"
)

func HandleBooks(resp http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		listBooks(resp)
	case http.MethodPost:
		addBook(req, resp)
	default:
		unsupportedHttpMethod(resp)
	}
}

func HandleBook(resp http.ResponseWriter, req *http.Request) {
	isbn := strings.TrimPrefix(req.URL.Path, "/api/book/")

	switch req.Method {
	case http.MethodGet:
		getBook(isbn, resp)
	case http.MethodPut:
		updateBook(isbn, resp)
	case http.MethodDelete:
		deleteBook(isbn, resp)
	default:
		unsupportedHttpMethod(resp)
	}
}

func unsupportedHttpMethod(resp http.ResponseWriter) {
	_, _ = resp.Write([]byte("unsupported http method"))
	resp.WriteHeader(http.StatusBadRequest)
}
