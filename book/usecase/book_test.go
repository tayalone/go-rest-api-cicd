package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tayalone/go-rest-api-cicd/book/entity"

	_bookMock "github.com/tayalone/go-rest-api-cicd/book/mock"
	"github.com/tayalone/go-rest-api-cicd/book/usecase"
)

func TestGetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockBookRepo := new(_bookMock.MockRepository)

		expectBook := entity.BookEntity{
			ID:      1,
			Title:   "Book no 1",
			SoldOut: true,
		}

		var id uint = 1

		mockBookRepo.On("GetByID", id).Return(expectBook, nil).Once()

		u := usecase.Initialize(mockBookRepo)

		book, err := u.GetByID(id)

		assert.NoError(t, err)
		assert.Equal(t, expectBook, book)
		mockBookRepo.AssertExpectations(t)
	})
	t.Run("errors: book not false", func(t *testing.T) {
		mockBookRepo := new(_bookMock.MockRepository)

		var id uint = 2

		expectBook := entity.BookEntity{}
		mockBookRepo.On("GetByID", id).Return(expectBook, errors.New("book not found")).Once()

		u := usecase.Initialize(mockBookRepo)

		_, err := u.GetByID(id)
		assert.Error(t, err)

		mockBookRepo.AssertExpectations(t)
	})
}

func TestCheckBookAvailableByID(t *testing.T) {
	t.Run("success: available", func(t *testing.T) {
		mockBookRepo := new(_bookMock.MockRepository)

		expectBook := entity.BookEntity{
			ID:      1,
			Title:   "Book Available",
			SoldOut: false,
		}

		var id uint = 1

		mockBookRepo.On("GetByID", id).Return(expectBook, nil).Once()

		u := usecase.Initialize(mockBookRepo)

		found, err := u.CheckBookAvailableByID(id)

		assert.Equal(t, found, true)
		assert.NoError(t, err)

		mockBookRepo.AssertExpectations(t)

	})
	t.Run("success: not available", func(t *testing.T) {
		mockBookRepo := new(_bookMock.MockRepository)

		expectBook := entity.BookEntity{
			ID:      2,
			Title:   "Book Not Available",
			SoldOut: true,
		}

		var id uint = 2

		mockBookRepo.On("GetByID", id).Return(expectBook, nil).Once()

		u := usecase.Initialize(mockBookRepo)

		found, err := u.CheckBookAvailableByID(id)

		assert.Equal(t, found, false)
		assert.NoError(t, err)

		mockBookRepo.AssertExpectations(t)
	})

	t.Run("errors: book not false", func(t *testing.T) {
		mockBookRepo := new(_bookMock.MockRepository)

		expectBook := entity.BookEntity{}

		var id uint = 1

		mockBookRepo.On("GetByID", id).Return(expectBook, errors.New("book not found")).Once()

		u := usecase.Initialize(mockBookRepo)

		found, err := u.CheckBookAvailableByID(id)

		assert.Equal(t, found, false)
		assert.Error(t, err)

		mockBookRepo.AssertExpectations(t)
	})
}
