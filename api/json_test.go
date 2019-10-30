package api

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_bookToJson(t *testing.T) {

	book := &Book{
		Title:       "My book",
		ISBN:        "123456",
		Author:      "me",
		Description: "",
	}
	json, err := bookToJson(book)

	assert.Nil(t, err, "error should be nil")

	assert.Equal(t,
		`{"title":"My book","isbn":"123456","author":"me","description":""}`,
		string(json))
}

func Test_booksToJson(t *testing.T) {

	json, err := booksToJson([]*Book{
		{
			Title:       "My book",
			ISBN:        "123456",
			Author:      "me",
			Description: "",
		},
		{
			Title:       "My book2",
			ISBN:        "123456",
			Author:      "him",
			Description: "",
		},
	})

	assert.Nil(t, err, "error should be nil")

	expected := strings.Join([]string{
		`[{"title":"My book","isbn":"123456","author":"me","description":""},`,
		`{"title":"My book2","isbn":"123456","author":"him","description":""}]`,
	}, "")

	assert.Equal(t, expected, string(json))
}

func Test_jsonToBook(t *testing.T) {
	jsonBytes := []byte(`{"title":"My book","isbn":"123456","author":"me","description":""}`)

	book, err := jsonToBook(jsonBytes)
	assert.Nil(t, err, "error should be nil")

	assert.Equal(t, &Book{
		Title:       "My book",
		ISBN:        "123456",
		Author:      "me",
		Description: "",
	}, book)
}
