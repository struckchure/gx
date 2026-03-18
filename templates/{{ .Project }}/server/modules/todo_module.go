package modules

import (
	"net/http"

	"github.com/a-h/rest"
	"github.com/labstack/echo/v5"
	"github.com/samber/do"
	"github.com/struckchure/gx/adapters"
	"github.com/struckchure/gx_app/dto"
	"github.com/struckchure/gx_app/ent"
	"github.com/struckchure/gx_app/handlers"
	"github.com/struckchure/gx_app/internals"
	"github.com/struckchure/gx_app/services"
)

func TodoModule(i *do.Injector) {
	do.Provide(i, handlers.NewTodoHandler)
	do.Provide(i, services.NewTodoService)

	srv := do.MustInvoke[*echo.Echo](i)
	todoHandler := do.MustInvoke[handlers.ITodoHandler](i)

	adapters.EchoV5R(srv.GET("/todo", todoHandler.List)).
		HasRequestModel(rest.ModelOf[dto.TodoListDto]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[[]ent.Todo]())

	adapters.EchoV5R(srv.POST("/todo", todoHandler.Create)).
		HasRequestModel(rest.ModelOf[dto.TodoCreateDto]()).
		HasResponseModel(http.StatusCreated, rest.ModelOf[internals.Empty]())

	adapters.EchoV5R(srv.PATCH("/todo/:id", todoHandler.Update)).
		HasRequestModel(rest.ModelOf[dto.TodoUpdateDto]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[internals.Empty]())

	adapters.EchoV5R(srv.DELETE("/todo/:id", todoHandler.Delete)).
		HasRequestModel(rest.ModelOf[dto.TodoDeleteDto]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[[]ent.Todo]())
}
