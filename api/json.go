package api

import "encoding/json"

func bookToJson(book *Book) ([]byte, error) {
	bytes, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func booksToJson(books []*Book) ([]byte, error) {
	bytes, err := json.Marshal(books)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func jsonToBook(jsonBytes []byte) (*Book, error) {
	book := &Book{}
	err := json.Unmarshal(jsonBytes, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
