package model

import "fmt"

//Books - the list of available books
var Books BookList

type BookDto struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	NoPages     int `json:"noPages"`
	ReleaseDate string `json:"releaseDate"`
	Author      AuthorDto `json:"author"`
}

type BookList []BookDto

func (b BookDto) String() string {
	return fmt.Sprintf("BookDto{UUID=%s, Title=%s, NoPages=%d, ReleaseDate=%s, Author=%s}",
		b.UUID, b.Title, b.NoPages, b.ReleaseDate, b.Author)
}

func (b *BookList) Get(bookUUID string) (BookDto, error) {
	err := fmt.Errorf("Could not find book by UUID %s", bookUUID)
	for _, book := range *b {
		if book.UUID == bookUUID {
			return book, nil
		}
	}
	return BookDto{}, err
}

func (b *BookList) Add(book BookDto) {
	*b = append(*b, book)
}

func (b *BookList) Update(updatedBook BookDto) (BookDto, error) {
	err := fmt.Errorf("Could not find book by UUID %s", updatedBook.UUID)
	var newBooks BookList
	for _, book := range *b {
		if book.UUID == updatedBook.UUID {
			newBooks = append(newBooks, updatedBook)
			err = nil
		} else {
			newBooks = append(newBooks, book)
		}
	}
	if err == nil {
		*b = newBooks
	}
	return updatedBook, err
}

func (b *BookList) Delete(bookUUID string) error {
	err := fmt.Errorf("Could not find book by UUID %s", bookUUID)
	var updatedBooks BookList
	for _, book := range *b {
		if book.UUID == bookUUID {
			err = nil
		} else {
			updatedBooks = append(updatedBooks, book)
		}
	}
	if err == nil {
		*b = updatedBooks
	}
	return err
}