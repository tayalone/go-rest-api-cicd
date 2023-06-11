package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/go-rest-api-cicd/book/port"
	// BooksRoute "github.com/tayalone/go-rest-api-cicd-private/api/books"
	// "github.com/tayalone/go-rest-api-cicd-private/book/port"
)

// API interface
type API interface {
	Start()
}

type api struct {
	*gin.Engine
}

func Initialize(bookUseCase port.Usecase) API {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// BooksRoute.Setup(r, bookUseCase)

	newAPI := &api{
		Engine: r,
	}
	return newAPI

}

func (a *api) Start() {
	str := ":" + strconv.FormatUint(uint64(8081), 10)

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the HTTP server in a goroutine
	srv := &http.Server{
		Addr:    str,
		Handler: a.Engine,
	}

	// Start http server in go routine
	go func() {
		if err := a.Engine.Run(str); err != nil && err != http.ErrServerClosed {
			log.Println("Failed to Start API Server", "errMsg", err.Error())
		}
	}()

	// Wait for an interrupt signal
	sig := <-quit
	log.Println("Server's shutting down", "sigName", sig.String())

	// Create a context with a timeout of 20 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Shut down the HTTP server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown", "errMsg", err.Error())
	}

	log.Println("API Server Stop")

}
