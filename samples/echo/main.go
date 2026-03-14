package main

import (
	"log"
	"net/http"

	"github.com/a-h/rest"
	"github.com/a-h/rest/swaggerui"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v5"
	"github.com/struckchure/gx"
	"github.com/struckchure/gx/adapters"
	"github.com/struckchure/gx/samples/echo/dto"
	"github.com/struckchure/gx/samples/echo/types"
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

	gx.GxSetup("GX + Echo")

	adapters.EchoV5R(e.GET("/", func(c *echo.Context) error {
		return c.String(200, "Hello")
	})).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]())

	adapters.EchoV5R(e.POST("/:three", func(c *echo.Context) error {
		return c.JSON(200, Response{Ok: true})
	})).
		HasRequestModel(rest.ModelOf[Request]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[[]Response]()).
		HasTags([]string{"numbers"})

	adapters.EchoV5R(e.POST("/:three/4", func(c *echo.Context) error {
		return c.JSON(200, Response{Ok: true})
	})).
		HasRequestModel(rest.ModelOf[dto.Request]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[dto.Response]()).
		HasTags([]string{"numbers"})

	adapters.EchoV5R(e.POST("/:three/5", func(c *echo.Context) error {
		return c.JSON(200, Response{Ok: true})
	})).
		HasRequestModel(rest.ModelOf[Request]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[types.Response]()).
		HasTags([]string{"numbers"})

	spec := gx.GxGenerate(func(spec *openapi3.T) {
		spec.Info.Version = "v1.0.0"
		spec.Info.Description = "This was generated using Gx (powered by <a href='https://github.com/a-h/rest'>a-h/rest</a>) and Echo v5"
	})

	// Attach the UI handler.
	ui, err := swaggerui.New(spec)
	if err != nil {
		log.Fatalf("failed to create swagger UI handler: %v", err)
	}
	e.GET("/swagger-ui*", echo.WrapHandler(ui))

	e.Start(":8000")
}
