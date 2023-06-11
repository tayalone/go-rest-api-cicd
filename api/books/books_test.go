package books_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tayalone/go-rest-api-cicd/api/books"
	"github.com/tayalone/go-rest-api-cicd/book/entity"
	_bookMock "github.com/tayalone/go-rest-api-cicd/book/mock"
)

func TestBookHandlers(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// create mock book usecase
	mockBookUsecase := new(_bookMock.MockUsecase)

	// Set up the routes
	books.Setup(router, mockBookUsecase)

	t.Run("GetBookByID", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			expectBook := entity.BookEntity{
				ID:      1,
				Title:   "Book no. 1",
				SoldOut: false,
			}
			id := uint(1)

			// Set up the mock bookUseCase response
			mockBookUsecase.On("GetByID", id).Return(expectBook, nil)
			// Create a new HTTP request to the appropriate endpoint
			req, _ := http.NewRequest("GET", "/books/"+"1", nil)
			w := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(w, req)

			// Assert the response status code
			assert.Equal(t, 200, w.Code)

			var response struct {
				Message string            `json:"message"`
				Book    entity.BookEntity `json:"book"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			// Assert the response body
			assert.Equal(t, "OK", response.Message)
			assert.Equal(t, expectBook, response.Book)

			mockBookUsecase.AssertExpectations(t)
		})
		t.Run("fail: `:id` is not uint", func(t *testing.T) {
			// Create a new HTTP request to the appropriate endpoint
			req, _ := http.NewRequest("GET", "/books/"+"one", nil)
			w := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(w, req)

			// Assert the response status code
			assert.Equal(t, http.StatusBadRequest, w.Code)
			var response struct {
				Message string `json:"message"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			// Assert the response body
			assert.Equal(t, "error must be uint", response.Message)
			mockBookUsecase.AssertExpectations(t)

		})
		t.Run("fail: not found book", func(t *testing.T) {
			expectBook := entity.BookEntity{}
			id := uint(2)

			// Set up the mock bookUseCase response
			mockBookUsecase.On("GetByID", id).Return(expectBook, errors.New("book not found"))
			// Create a new HTTP request to the appropriate endpoint
			req, _ := http.NewRequest("GET", "/books/"+"2", nil)
			w := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(w, req)

			// Assert the response status code
			assert.Equal(t, http.StatusNotFound, w.Code)
			var response struct {
				Message string `json:"message"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			// Assert the response body
			assert.Equal(t, "book not found", response.Message)
			mockBookUsecase.AssertExpectations(t)

		})
	})

	t.Run("CheckBookAvailabilityByID", func(t *testing.T) {
		t.Run("success: book available", func(t *testing.T) {

			expectAvailable := true
			var expectError error = nil

			id := uint(1)

			// Set up the mock bookUseCase response
			mockBookUsecase.On("CheckBookAvailableByID", id).Return(expectAvailable, expectError)
			// Create a new HTTP request to the appropriate endpoint
			req, _ := http.NewRequest("GET", "/books/"+"1"+"/available", nil)
			w := httptest.NewRecorder()
			// Serve the request using the router
			router.ServeHTTP(w, req)

			// Assert the response status code
			assert.Equal(t, 200, w.Code)

			var response struct {
				Message   string `json:"message"`
				Available bool   `json:"available"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			// Assert the response body
			assert.Equal(t, "OK", response.Message)
			assert.Equal(t, expectAvailable, true)

			mockBookUsecase.AssertExpectations(t)

		})

		t.Run("success: book not available", func(t *testing.T) {
			expectAvailable := false
			var expectError error = nil
			id := uint(2)

			// Set up the mock bookUseCase response
			mockBookUsecase.On("CheckBookAvailableByID", id).Return(expectAvailable, expectError)
			// Create a new HTTP request to the appropriate endpoint
			req, _ := http.NewRequest("GET", "/books/"+"2"+"/available", nil)
			w := httptest.NewRecorder()
			// Serve the request using the router
			router.ServeHTTP(w, req)

			// Assert the response status code
			assert.Equal(t, 200, w.Code)

			var response struct {
				Message   string `json:"message"`
				Available bool   `json:"available"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			// Assert the response body
			assert.Equal(t, "OK", response.Message)
			assert.Equal(t, expectAvailable, false)

			mockBookUsecase.AssertExpectations(t)
		})
		t.Run("fail: `:id` is not uint", func(t *testing.T) {
			// Create a new HTTP request to the appropriate endpoint
			// expectAvailable := true
			expectError := errors.New("error must be uint")

			req, _ := http.NewRequest("GET", "/books/"+"one"+"/available", nil)
			w := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(w, req)

			// Assert the response status code
			assert.Equal(t, http.StatusBadRequest, w.Code)
			var response struct {
				Message string `json:"message"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			// Assert the response body
			assert.Equal(t, expectError.Error(), response.Message)
			mockBookUsecase.AssertExpectations(t)
		})
		t.Run("fail: not found book", func(t *testing.T) {
			id := uint(3)

			expectAvailable := false
			expectError := errors.New("book not found")

			// Set up the mock bookUseCase response
			mockBookUsecase.On("CheckBookAvailableByID", id).Return(expectAvailable, expectError)
			// Create a new HTTP request to the appropriate endpoint
			req, _ := http.NewRequest("GET", "/books/"+"3"+"/available", nil)
			w := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(w, req)

			// Assert the response status code
			assert.Equal(t, http.StatusNotFound, w.Code)
			var response struct {
				Message string `json:"message"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			// Assert the response body
			assert.Equal(t, expectError.Error(), response.Message)
			mockBookUsecase.AssertExpectations(t)
		})
	})

}
