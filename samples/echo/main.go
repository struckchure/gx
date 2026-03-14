package main

import (
	"log"
	"net/http"

	"github.com/a-h/rest"
	"github.com/a-h/rest/swaggerui"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v5"
	"github.com/struckchure/gx"
)

type Request struct {
	One   string `json:"one"`
	Two   string `query:"two"`
	Three string `param:"three"`
}

type Response struct {
	Ok bool `json:"ok"`
}

func main() {
	e := echo.New()

	gx.Setup("messages")

	gx.Doc[gx.None](e.GET("/", func(c *echo.Context) error {
		return c.String(200, "Hello")
	})).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]())

	gx.Doc[Request](e.POST("/:three", func(c *echo.Context) error {
		return c.JSON(200, Response{Ok: true})
	})).
		HasResponseModel(http.StatusOK, rest.ModelOf[Response]()).
		HasTags([]string{"numbers"})

	spec := gx.Generate(func(spec *openapi3.T) {
		spec.Info.Version = "v1.0.0"
		spec.Info.Description = "Messages API"
	})

	// Attach the UI handler.
	ui, err := swaggerui.New(spec)
	if err != nil {
		log.Fatalf("failed to create swagger UI handler: %v", err)
	}
	e.GET("/swagger-ui*", echo.WrapHandler(ui))

	e.Start(":8000")
}
