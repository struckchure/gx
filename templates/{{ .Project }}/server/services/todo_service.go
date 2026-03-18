package services

import (
	"context"

	"github.com/samber/do"
	"github.com/struckchure/gx_app/dto"
	"github.com/struckchure/gx_app/ent"
	"github.com/struckchure/gx_app/ent/todo"
)

type ITodoService interface {
	List(dto.TodoListDto) ([]*ent.Todo, error)
	Create(dto.TodoCreateDto) error
	Update(dto.TodoUpdateDto) error
	Delete(dto.TodoDeleteDto) error
}

type TodoService struct {
	db *ent.Client
}

func (t *TodoService) List(todoListDto dto.TodoListDto) ([]*ent.Todo, error) {
	return t.db.Todo.Query().
		Where(todo.StatusEQ(todoListDto.Status)).
		All(context.Background())
}

func (t *TodoService) Create(todoCreateDto dto.TodoCreateDto) error {
	return t.db.Todo.Create().
		SetTitle(todoCreateDto.Title).
		SetDescription(todoCreateDto.Description).
		Exec(context.Background())
}

func (t *TodoService) Update(todoUpdateDto dto.TodoUpdateDto) error {
	return t.db.Todo.Update().
		Where(todo.ID(todoUpdateDto.Id)).
		SetNillableTitle(&todoUpdateDto.Title).
		SetNillableDescription(&todoUpdateDto.Description).
		SetNillableStatus(&todoUpdateDto.Status).
		Exec(context.Background())
}

func (t *TodoService) Delete(todoDeleteDto dto.TodoDeleteDto) error {
	return t.db.Todo.DeleteOneID(todoDeleteDto.Id).Exec(context.Background())
}

func NewTodoService(i *do.Injector) (ITodoService, error) {
	return &TodoService{db: do.MustInvoke[*ent.Client](i)}, nil
}
