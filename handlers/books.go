// Package classification Book API
//
// Documentation for Book API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"monitoring/data"
	"net/http"
)

// just for the sake of documentation with swagger.
// responses:
//   200: booksResponse

// swagger:response booksResponse
type booksResposeWrapper struct {
	// in:body
	Body []data.Book
}

// a single book
// swagger:response bookResponse
type bookResponseWrapper struct {
	// in: body
	Body data.Book
}

// swagger:response noContentResponse
type bookNoContent struct {
}

// swagger:parameters deleteBook
type bookIDParameterWrapper struct {
	// The id of the book to delete from the books database
	// in:path
	// required:true
	ID int `json:"id"`
}

// Books is a http.Handler
type Books struct {
	l *log.Logger
}

// NewBooks creates a books handler.
func NewBooks(l *log.Logger) *Books {
	return &Books{l}
}

// Note all the below functions should be exported.

// KeyBook is used because we would need a key for context and the recommended type is struct.
type KeyBook struct{}

// MiddlewareBooksValidation used to validate the book in the request and calls next if everything is ok.
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
