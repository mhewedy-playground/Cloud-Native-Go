package api

type Book struct {
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
