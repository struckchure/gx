package adapters

import (
	"github.com/a-h/rest"
)

var api = rest.NewAPI("messages")

func init() {
	api.StripPkgPaths = []string{"github.com/a-h/rest/example", "github.com/a-h/respond"}
	// api.RegisterModel(rest.ModelOf[respond.Error](), rest.WithDescription("Standard JSON error"), func(s *openapi3.Schema) {
	// 	status := s.Properties["statusCode"]
	// 	status.Value.WithMin(100).WithMax(600)
	// })
}

type Route[C any] struct{}

func (r *Route[C]) Request(model any) *Route[C] {
	return r
}

func (r *Route[C]) Response(status int, model any) *Route[C] {
	return r
}

func (r *Route[C]) Handle(method string) (string, func(C) error) {
	return method, func(c C) error { return nil }
}

type IRoute[C any] interface {
	Handle(method string) (string, func(C) error)
	Request(model any) *Route[C]
	Response(status int, model any)
}
