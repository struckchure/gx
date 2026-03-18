//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Hooks: []gen.Hook{
			RemoveOmitEmptyIfNillableStructTag("json"),
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func RemoveOmitEmptyIfNillableStructTag(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					if field.Nillable {
						field.StructTag = fmt.Sprintf(`json:"%s"`, field.Name)
					}
				}
			}
			return next.Generate(g)
		})
	}
}
