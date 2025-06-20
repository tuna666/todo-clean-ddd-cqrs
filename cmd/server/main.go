package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tuna666/todo-clean-ddd-cqrs/internal/adapter/handler"
	adapterRepo "github.com/tuna666/todo-clean-ddd-cqrs/internal/adapter/repository"
	"github.com/tuna666/todo-clean-ddd-cqrs/internal/usecase"
	"github.com/tuna666/todo-clean-ddd-cqrs/pkg/database"
)

func main() {
	database.Init()

	// DI
	repo := adapterRepo.NewTodoGorm()
	cmdSvc, querySvc := usecase.NewTodoUseCase(repo)
	h := handler.NewTodoHandler(cmdSvc, querySvc)

	// Echo
	e := echo.New()
	e.POST("/todos", h.Create)
	e.GET("/todos", h.List)

	e.Logger.Fatal(e.Start(":8080"))
}