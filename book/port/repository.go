package port

import "github.com/tayalone/go-rest-api-cicd/book/entity"

type Usecase interface {
	GetByID(id uint) (entity.BookEntity, error)
	CheckBookAvailableByID(id uint) (bool, error)
}
