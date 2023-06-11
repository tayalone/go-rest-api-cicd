package books

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/go-rest-api-cicd/book/port"
)

func Setup(r *gin.Engine, bookUseCase port.Usecase) {

	book := r.Group("/books/:id")
	{
		book.GET("", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "error must be uint",
				})
				return
			}

			book, err := bookUseCase.GetByID(uint(id))

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"book":    book,
			})
		})

		book.GET("/available", func(c *gin.Context) {

			idStr := c.Param("id")
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "error must be uint",
				})
				return
			}

			available, err := bookUseCase.CheckBookAvailableByID(uint(id))

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":   "OK",
				"available": available,
			})
		})
	}

}
