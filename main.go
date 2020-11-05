package main

import (
	"log"
	"monitoring/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rusart/muxprom"
)

var prom *muxprom.MuxProm

func main() {

	l := log.New(os.Stdout, "books-api", log.LstdFlags)

	r := mux.NewRouter()
	prom = muxprom.New(
		muxprom.Router(r),
		//muxprom.MetricsRouteName("prommetrics"),
		muxprom.MetricsPath("/metrics"),
	)
	prom.Instrument()

	bh := handlers.NewBooks(l)

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/books", bh.GetBooks)
	getRouter.HandleFunc("/api/books/{id:[0-9]+}", bh.GetBook)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/api/books/{id:[0-9]+}", bh.UpdateBook)
	putRouter.Use(bh.MiddlewareBooksValidation)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/books", bh.AddBook)
	postRouter.Use(bh.MiddlewareBooksValidation)

	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/api/books/{id:[0-9]+}", bh.DeleteBook)

	l.Fatal(http.ListenAndServe(":8000", r))
}
