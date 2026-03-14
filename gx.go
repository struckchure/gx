package gx

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/a-h/rest"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v5"
)

var api *rest.API

func Setup(name string, opts ...rest.APIOpts) *rest.API {
	if api != nil {
		return api
	}

	api = rest.NewAPI(name, opts...)

	return api
}

func errorIfNotApi() {
	if api == nil {
		log.Fatalln("Setup method must be called before defining models")
	}
}

func Generate(mods ...func(*openapi3.T)) *openapi3.T {
	spec, err := api.Spec()
	if err != nil {
		log.Fatalf("failed to create spec: %v", err)
	}

	for _, mod := range mods {
		mod(spec)
	}

	return spec
}

func HasRequestModel[T any](route *rest.Route) *rest.Route {
	errorIfNotApi()

	t := reflect.TypeFor[T]()
	if t.Kind() != reflect.Struct {
		return route
	}

	var fields []reflect.StructField
	for i := range t.NumField() {
		field := t.Field(i)
		isParam := false

		for tag := range strings.SplitSeq(string(field.Tag), " ") {
			tag = strings.TrimSpace(tag)
			if tag == "" {
				continue
			}
			parts := strings.SplitN(tag, ":", 2)
			if len(parts) != 2 {
				continue
			}
			key := parts[0]
			value := strings.Trim(parts[1], `"`)
			switch key {
			case "param":
				route = route.HasPathParameter(value, rest.PathParam{})
				isParam = true
			case "query":
				route = route.HasQueryParameter(value, rest.QueryParam{})
				isParam = true
			}
		}

		if !isParam {
			fields = append(fields, field)
		}
	}

	modelType := reflect.StructOf(fields)

	return route.HasRequestModel(rest.Model{Type: modelType})
}

func formatToOpenApi(path string) string {
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

type None any

func Doc[T any](info echo.RouteInfo) *rest.Route {
	errorIfNotApi()

	route := api.Route(info.Method, formatToOpenApi(info.Path))

	return HasRequestModel[T](route)
}
