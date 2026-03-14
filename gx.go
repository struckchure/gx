package gx

import (
	"log"
	"strings"

	"github.com/a-h/rest"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/samber/lo"
)

func GxSetup(name string, opts ...rest.APIOpts) *rest.API {
	if Api != nil {
		return Api
	}

	opts = append(opts, func(a *rest.API) {
		a.ApplyPostNormalizeTransform = func(normalized string) string {
			normalized = strings.Join(last(strings.Split(normalized, "_"), 2), "_")
			return lo.PascalCase(normalized)
		}
	})
	Api = rest.NewAPI(name, opts...)

	return Api
}

func GxGenerate(mods ...func(*openapi3.T)) *openapi3.T {
	spec, err := Api.Spec()
	if err != nil {
		log.Fatalf("failed to create spec: %v", err)
	}

	for _, mod := range mods {
		mod(spec)
	}

	return spec
}
