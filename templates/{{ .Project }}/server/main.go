package main

import (
	"github.com/samber/do"
	"github.com/struckchure/gx"
	"github.com/struckchure/gx_app/internals"
	"github.com/struckchure/gx_app/modules"
)

func main() {
	// Create the dependency injector
	i := do.New()

	// Setup GX API first (before registering routes)
	gx.GxSetup("Todo API")

	// Register modules with the injector
	modules.RootModule(i)
	modules.TodoModule(i)

	// Start the server with the injector
	internals.StartServer(i)
}
