package usecase

import (
	"github.com/tayalone/go-rest-api-cicd/book/entity"
	"github.com/tayalone/go-rest-api-cicd/book/port"
)

type book struct {
	repo port.Repository
}

func Initialize(repo port.Repository) port.Usecase {
	return &book{
		repo: repo,
	}
}

func (b *book) GetByID(id uint) (entity.BookEntity, error) {
	return b.repo.GetByID(id)
}

func (b *book) CheckBookAvailableByID(id uint) (bool, error) {
	book, err := b.repo.GetByID(id)
	if err != nil {
		return false, err
	}

	return book.IsAvailable(), nil
}
