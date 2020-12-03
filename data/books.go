package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
)

// Book defines the structure for an API Book
// isbn can be 10 or 13 digits long.
// swagger:model
type Book struct {
	// the id of the book
	// min:1
	ID int `json:"id"`

	// isbn code of the book
	// required:true
	// min:10
	Isbn string `json:"isbn" validate:"required,isbn"`

	// title of the book
	Title string `json:"title"`

	// author of the book
	Author *Author `json:"author"`
}

// Author Structure
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Books []*Book

// Validate method implementing Secure Coding Principles.
func (b *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}

func (b *Book) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(b)
}

func (b *Books) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(b)
}

func (b *Book) ToJSONbook(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(b)
}

func getNextID() int {
	lb := booksList[len(booksList)-1]
	return lb.ID + 1
}

var ErrorBookNotFound = fmt.Errorf("Book not found")

func findBook(id int) (*Book, int, error) {
	for i, b := range booksList {
		if b.ID == id {
			return b, i, nil
		}
	}

	return nil, -1, ErrorBookNotFound
}

func GetBooks() Books {
	return booksList
}

func GetBook(id int) (*Book, error) {
	bk, _, err := findBook(id)
	if err != nil {
		return nil, err
	}

	return bk, nil
}

func AddBook(b *Book) {
	if b.ID == 0 {
		b.ID = getNextID()
	}
	booksList = append(booksList, b)
}

func UpdateBook(b *Book, id int) error {
	_, pos, err := findBook(id)
	if err != nil {
		return err
	}

	b.ID = id
	booksList[pos] = b

	return nil
}

func DeleteBook(id int) error {
	_, pos, err := findBook(id)
	if err != nil {
		return err
	}

	booksList = append(booksList[:pos], booksList[pos+1:]...)

	return nil
}

// Init books var as a slice Book struct
var booksList = []*Book{
	&Book{
		ID:     1,
		Isbn:   "978-0812036381",
		Title:  "Hamlet",
		Author: &Author{Firstname: "William", Lastname: "Shakespeare"},
	},
	&Book{
		ID:     2,
		Isbn:   "978-0671027032",
		Title:  "How to Win Friends & Influence People",
		Author: &Author{Firstname: "Dale", Lastname: "Carnegie"},
	},
}
