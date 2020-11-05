package data

import "testing"

func TestCehckValidation(t *testing.T) {
	b := &Book{
		Isbn: "978-1544507859",
	}

	err := b.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
