package internals

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/samber/do"
	"github.com/struckchure/gx_app/ent"
)

func NewDatabase(i *do.Injector) (*ent.Client, error) {
	env := do.MustInvoke[*Env](i)

	// Create an ent.Client by connecting to the database.
	client, err := ent.Open("postgres", env.DATABASE_URL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return client, nil
}
