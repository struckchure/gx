package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/samber/do"
	"github.com/struckchure/gx_app/dto"
	"github.com/struckchure/gx_app/services"
)

type ITodoHandler interface {
	List(*echo.Context) error
	Create(*echo.Context) error
	Update(*echo.Context) error
	Delete(*echo.Context) error
}

type TodoHandler struct {
	todoService services.ITodoService
}

func (t *TodoHandler) List(c *echo.Context) error {
	var request dto.TodoListDto
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := t.todoService.List(request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (t *TodoHandler) Create(c *echo.Context) error {
	var request dto.TodoCreateDto
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := t.todoService.Create(request)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (t *TodoHandler) Update(c *echo.Context) error {
	var request dto.TodoUpdateDto
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := t.todoService.Update(request)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusAccepted)
}

func (t *TodoHandler) Delete(c *echo.Context) error {
	var request dto.TodoDeleteDto
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := t.todoService.Delete(request)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func NewTodoHandler(i *do.Injector) (ITodoHandler, error) {
	return &TodoHandler{
		todoService: do.MustInvoke[services.ITodoService](i),
	}, nil
}
