package dto

import "github.com/struckchure/gx_app/ent/todo"

type TodoListDto struct {
	Status todo.Status `query:"status"`
}

type TodoCreateDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoUpdateDto struct {
	Id          int         `param:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      todo.Status `json:"status"`
}

type TodoDeleteDto struct {
	Id int `param:"id"`
}
