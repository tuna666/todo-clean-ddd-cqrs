package repository

import (
	"github.com/tuna666/todo-clean-ddd-cqrs/internal/domain/todo"
	"github.com/tuna666/todo-clean-ddd-cqrs/pkg/database"
)

type todoGorm struct{}

func NewTodoGorm() todo.Repository {
	// マイグレーションもここで簡易実行
	database.DB.AutoMigrate(&todo.Todo{})
	return &todoGorm{}
}

func (r *todoGorm) Create(t *todo.Todo) error {
	return database.DB.Create(t).Error
}

func (r *todoGorm) List() ([]todo.Todo, error) {
	var out []todo.Todo
	err := database.DB.Find(&out).Error
	return out, err
}
