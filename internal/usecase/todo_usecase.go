package usecase

import "github.com/tuna666/todo-clean-ddd-cqrs/domain/todo"

// Command 層
type CreateTodoCommand struct {
	Title string `json:"title"`
}

type TodoCommandService interface {
	Create(cmd CreateTodoCommand) (todo.ID, error)
}

// Query 層
type TodoQueryService interface {
	List() ([]todo.Todo, error)
}

// 1 構造体で両方満たす
type todoUC struct {
	repo todo.Repository
}

func NewTodoUseCase(r todo.Repository) (TodoCommandService, TodoQueryService) {
	uc := &todoUC{repo: r}
	return uc, uc
}

// --- Command ---
func (uc *todoUC) Create(cmd CreateTodoCommand) (todo.ID, error) {
	t := &todo.Todo{Title: cmd.Title}
	if err := uc.repo.Create(t); err != nil {
		return 0, err
	}
	return t.ID, nil
}

// --- Query ---
func (uc *todoUC) List() ([]todo.Todo, error) {
	return uc.repo.List()
}
