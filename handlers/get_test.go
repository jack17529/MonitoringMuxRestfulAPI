package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func SubRouter() *mux.Router {
	r := mux.NewRouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()
	l := log.New(os.Stdout, "books-api", log.LstdFlags)
	bh := NewBooks(l)
	getRouter.HandleFunc("/api/books", bh.GetBooks)
	return getRouter
}

func TestGetBooks(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/books", nil)
	response := httptest.NewRecorder()
	SubRouter().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", response.Code)
	}
	assert.Equal(t, 200, response.Code, "OK response is expected.")
}
