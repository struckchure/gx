package adapters

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/struckchure/gx"
)

func formatEchoUrlToOpenApi(path string) string {
	if path == "/" {
		return path
	}

	split := strings.Split(path, "/")
	result := make([]string, len(split))

	for i, s := range split {
		if after, ok := strings.CutPrefix(s, ":"); ok {
			result[i] = fmt.Sprintf("{%s}", after)
		} else {
			result[i] = s
		}
	}

	return strings.Join(result, "/")
}

func EchoV5R(info echo.RouteInfo) *gx.GxRoute {
	gx.ErrorIfNotApi()

	route := gx.Api.Route(info.Method, formatEchoUrlToOpenApi(info.Path))

	return &gx.GxRoute{Route: route}
}
