package main

import (
	"fmt"

	"github.com/tayalone/go-rest-api-cicd/api"
	"github.com/tayalone/go-rest-api-cicd/book/repository"
	"github.com/tayalone/go-rest-api-cicd/book/usecase"
)

func main() {
	fmt.Println("let's start")

	fmt.Println("hello world !!")

	bookRepo := repository.Initialize()
	bookUseCase := usecase.Initialize(bookRepo)

	myAPI := api.Initialize(bookUseCase)

	myAPI.Start()
}
