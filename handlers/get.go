package handlers

import (
	"monitoring/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route GET /api/books books listBooks
// Returns a list of books from the database
// responses:
//	200: booksResponse

// GetBooks gets All Books
func (b *Books) GetBooks(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle GET requests")

	lb := data.GetBooks()

	// instead of marshalling we would use encoder.
	err := lb.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /api/books/{id} books listSingleBook
// Return a list of books from the database
// responses:
//	200: bookResponse

// GetBook gets a singe book.
func (b *Books) GetBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle GET requests")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Unable to convert id from string to integer", http.StatusBadRequest)
		return
	}

	bk, err := data.GetBook(id)
	if err == data.ErrorBookNotFound {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Book not found", http.StatusInternalServerError)
		return
	}

	// instead of marshalling we would use encoder.
	err = bk.ToJSONbook(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}

}
