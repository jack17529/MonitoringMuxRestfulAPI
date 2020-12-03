package handlers

import (
	"monitoring/data"
	"net/http"
)

// swagger:route POST /api/books books saveBook
// Adds a new book
// responses:
//	200: booksResponse

// AddBook saves a single book in the database
func (b *Books) AddBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle POST requests")

	bk := r.Context().Value(KeyBook{}).(*data.Book)

	data.AddBook(bk)
}
