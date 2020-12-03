package handlers

import (
	"monitoring/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /api/books books updateBook
// Update a books details
//
// responses:
//	201: noContentResponse

// UpdateBook updates a book
func (b Books) UpdateBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle PUT requests")

	// params is just a map.
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Unable to convert id from string to integer", http.StatusBadRequest)
		return
	}

	bk := r.Context().Value(KeyBook{}).(*data.Book)

	err = data.UpdateBook(bk, id)
	if err == data.ErrorBookNotFound {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Book not found", http.StatusInternalServerError)
		return
	}
}
