package model

import "fmt"

//Books - the list of available books
var Books BookList

type Book struct {
	Entity
	Title       string `json:"title"`
	NoPages     int `json:"noPages"`
	ReleaseDate string `json:"releaseDate"`
	Author      Author `json:"author" gorm:"foreignkey:AuthorId"`
	AuthorId    int    `json:"-"`
}

type BookList []Book

func (b Book) String() string {
	return fmt.Sprintf("BookDto{UUID=%s, Title=%s, NoPages=%d, ReleaseDate=%s, Author=%s}",
		b.UUID, b.Title, b.NoPages, b.ReleaseDate, b.Author)
}

func (b *BookList) Get(bookUUID string) (Book, error) {
	err := fmt.Errorf("Could not find book by UUID %s", bookUUID)
	for _, book := range *b {
		if book.UUID == bookUUID {
			return book, nil
		}
	}
	return Book{}, err
}

func (b *BookList) Add(book Book) {
	*b = append(*b, book)
}

func (b *BookList) Update(updatedBook Book) (Book, error) {
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