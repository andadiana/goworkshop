package model

import "fmt"

//Authors - the list of available authors
var Authors AuthorList

type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

type AuthorList []AuthorDto

func (a AuthorDto) String() string {
	return fmt.Sprintf("AuthorDto{UUID=%s, FirstName=%s, LastName=%s, Birthday=%s, Death=%s}",
		a.UUID, a.FirstName, a.LastName, a.Birthday, a.Death)
}

func (a *AuthorList) Get(authorUUID string) (AuthorDto, error) {
	err := fmt.Errorf("Could not find author by UUID %s", authorUUID)
	for _, author := range *a {
		if author.UUID == authorUUID{
			return author, nil
		}
	}
	return AuthorDto{}, err
}

func (a *AuthorList) Add(author AuthorDto) {
	*a = append(*a, author)
}

func (a *AuthorList) Update(updatedAuthor AuthorDto) (AuthorDto, error) {
	err := fmt.Errorf("Could not find author by UUID %s", updatedAuthor.UUID)
	var newAuthors AuthorList
	for _, author := range *a {
		if author.UUID == updatedAuthor.UUID {
			newAuthors = append(newAuthors, updatedAuthor)
			err = nil
		} else {
			newAuthors = append(newAuthors, author)
		}
	}
	if err == nil {
		*a = newAuthors
	}
	return updatedAuthor, err
}

func (a *AuthorList) Delete(authorUUID string) error {
	err := fmt.Errorf("Could not find author by UUID %s", authorUUID)
	var updatedAuthors AuthorList
	for _, author := range *a {
		if author.UUID == authorUUID {
			err = nil
		} else {
			updatedAuthors = append(updatedAuthors, author)
		}
	}
	if err == nil {
		*a = updatedAuthors
	}
	return err
}