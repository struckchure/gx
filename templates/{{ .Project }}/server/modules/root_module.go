package modules

import (
	"github.com/samber/do"
	"github.com/struckchure/gx_app/internals"
)

func RootModule(i *do.Injector) {
	do.Provide(i, internals.NewEnv)
	do.Provide(i, internals.NewServer)
	do.Provide(i, internals.NewDatabase)
}
