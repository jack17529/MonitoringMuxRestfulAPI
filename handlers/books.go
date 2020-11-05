package handlers

import (
	"context"
	"fmt"
	"log"
	"monitoring/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Books struct {
	l *log.Logger
}

func NewBooks(l *log.Logger) *Books {
	return &Books{l}
}

// Note all the below functions should be exported.

// Get All Books
func (b *Books) GetBooks(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle GET requests")

	lb := data.GetBooks()

	// instead of marshalling we would use encoder.
	err := lb.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

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

func (b *Books) AddBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handle POST requests")

	bk := r.Context().Value(KeyBook{}).(*data.Book)

	data.AddBook(bk)
}

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

// we would need a key for context and the recommended type is struct.
type KeyBook struct{}

func (b *Books) MiddlewareBooksValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			bk := &data.Book{}

			err := bk.FromJSON(r.Body)
			if err != nil {
				b.l.Println("[ERROR] deserializing the book", err)
				http.Error(w, "Error reading the book", http.StatusBadRequest)
				return
			}

			// valide the book
			err = bk.Validate()
			if err != nil {
				b.l.Println("[ERROR] validating the book", err)
				http.Error(w, fmt.Sprintf("Error validating the book: %s", err), http.StatusBadRequest)
				return
			}

			// putting it in the request because it has the context.
			ctx := context.WithValue(r.Context(), KeyBook{}, bk)
			r = r.WithContext(ctx)

			// Calling the next handler, which can be another middleware in the chain,
			// or the final handler.
			next.ServeHTTP(w, r)
		})
}
