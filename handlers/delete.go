package handlers

import (
	"monitoring/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Deletes a single book.
func (b *Books) DeleteBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle DELETE requests")

	// params is just a map.
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Unable to convert id from string to integer", http.StatusBadRequest)
		return
	}

	// bk := &data.Book{}

	// // error here.
	// err = bk.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	// }

	err = data.DeleteBook(id)
	if err == data.ErrorBookNotFound {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Book not found", http.StatusInternalServerError)
		return
	}
}
