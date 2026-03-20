package adapters

import (
	"github.com/labstack/echo/v4"
	"github.com/struckchure/gx"
)

func EchoV4R(info *echo.Route) *gx.GxRoute {
	gx.ErrorIfNotApi()

	route := gx.Api.Route(info.Method, formatEchoUrlToOpenApi(info.Path))

	return &gx.GxRoute{Route: route}
}
