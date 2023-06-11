package repository

import (
	"errors"

	"github.com/tayalone/go-rest-api-cicd/book/entity"
	"github.com/tayalone/go-rest-api-cicd/book/port"
)

var BookStore = []entity.BookEntity{
	{
		ID:      1,
		Title:   "Book 1",
		SoldOut: false,
	},
	{
		ID:      2,
		Title:   "Book 2",
		SoldOut: true,
	},
}

type repo struct{}

func Initialize() port.Repository {
	return &repo{}
}

func (r *repo) GetByID(id uint) (entity.BookEntity, error) {
	for _, b := range BookStore {
		if b.ID == id {
			return b, nil
		}
	}
	return entity.BookEntity{}, errors.New("book not found")
}
