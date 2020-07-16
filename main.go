package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getAllBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	for _, item := range books{
		if params["id"]==item.ID{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID=strconv.Itoa(rand.Intn(10000000))	// Mock ID -not safe
	books = append(books,book)
	json.NewEncoder(w).Encode(&Book{})
}

func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if params["id"]==item.ID {
			books = append(books[:index], books[index+1:]...)
			//break
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID=params["id"]
			books = append(books,book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if params["id"]==item.ID {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main(){
	r := mux.NewRouter()
	
	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author : & Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "448744", Title: "Book Two", Author : & Author{Firstname: "Steve", Lastname: "Smith"}})
	
	r.HandleFunc("/api/books", getAllBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000",r))
}
