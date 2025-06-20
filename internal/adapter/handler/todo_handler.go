package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tuna666/todo-clean-ddd-cqrs/internal/usecase"
)

type TodoHandler struct {
	cmd usecase.TodoCommandService
	q   usecase.TodoQueryService
}

func NewTodoHandler(cmd usecase.TodoCommandService, q usecase.TodoQueryService) *TodoHandler {
	return &TodoHandler{cmd: cmd, q: q}
}

// POST /todos
func (h *TodoHandler) Create(c echo.Context) error {
	var req usecase.CreateTodoCommand
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	id, err := h.cmd.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"id": id})
}

// GET /todos
func (h *TodoHandler) List(c echo.Context) error {
	todos, err := h.q.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}
