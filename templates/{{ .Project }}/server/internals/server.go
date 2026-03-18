package internals

import (
	"fmt"
	"log"

	"github.com/a-h/rest/swaggerui"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/samber/do"
	"github.com/struckchure/gx"
)

func NewServer(i *do.Injector) (*echo.Echo, error) {
	e := echo.New()
	e.Logger = Logger

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			e.Logger.Info(fmt.Sprintf("[%s] %s %d - %s", v.Method, v.URI, v.Status, v.Latency))
			return nil
		},
	}))

	return e, nil
}

func StartServer(i *do.Injector) {
	srv := do.MustInvoke[*echo.Echo](i)

	spec := gx.GxGenerate(func(spec *openapi3.T) {
		spec.Info.Version = "v1.0.0"
		spec.Info.Description = "My Really Cool Todo API"
	})

	// Attach the UI handler.
	ui, err := swaggerui.New(spec)
	if err != nil {
		log.Fatalf("failed to create swagger UI handler: %v", err)
	}
	srv.GET("/swagger-ui*", echo.WrapHandler(ui))

	srv.Logger.Error(srv.Start(":8080").Error())
}
