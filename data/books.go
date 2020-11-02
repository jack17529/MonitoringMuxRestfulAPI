package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// Book Struct
type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Books []*Book

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
		Isbn:   "448743",
		Title:  "Book One",
		Author: &Author{Firstname: "John", Lastname: "Doe"},
	},
	&Book{
		ID:     2,
		Isbn:   "448744",
		Title:  "Book Two",
		Author: &Author{Firstname: "Steve", Lastname: "Smith"},
	},
}
