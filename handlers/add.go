package handlers

import (
	"monitoring/data"
	"net/http"
)

// Adds a single book.
func (b *Books) AddBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle POST requests")

	bk := r.Context().Value(KeyBook{}).(*data.Book)

	data.AddBook(bk)
}
