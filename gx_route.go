package gx

import (
	"reflect"
	"strings"

	"github.com/a-h/rest"
	"github.com/labstack/echo/v5"
)

type GxRoute struct {
	*rest.Route
}

func (g *GxRoute) HasRequestModel(request rest.Model) *GxRoute {
	ErrorIfNotApi()

	t := request.Type
	if t.Kind() != reflect.Struct {
		return g
	}

	for i := range t.NumField() {
		field := t.Field(i)

		for tag := range strings.SplitSeq(string(field.Tag), " ") {
			tag = strings.TrimSpace(tag)
			if tag == "" {
				continue
			}
			parts := strings.SplitN(tag, ":", 2)
			if len(parts) != 2 {
				continue
			}
			key, value := parts[0], strings.Trim(parts[1], `"`)
			switch key {
			case "param":
				g.Route = g.Route.HasPathParameter(value, rest.PathParam{})
			case "query":
				g.Route = g.Route.HasQueryParameter(value, rest.QueryParam{})
			}
		}
	}

	return g
}

type RArgs struct {
	Method string
	Path   string
}

func R(info echo.RouteInfo) *GxRoute {
	ErrorIfNotApi()

	route := Api.Route(info.Method, formatToOpenApi(info.Path))

	return &GxRoute{route}
}
