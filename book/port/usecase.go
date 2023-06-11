package port

import "github.com/tayalone/go-rest-api-cicd/book/entity"

type Repository interface {
	GetByID(id uint) (entity.BookEntity, error)
}
