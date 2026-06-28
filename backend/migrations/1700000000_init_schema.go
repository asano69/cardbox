package migrations

import (
	_ "embed"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

//go:embed schema.json
var schemaJSON []byte

func init() {
	m.Register(func(app core.App) error {
		return app.ImportCollectionsByMarshaledJSON(schemaJSON, false)
	}, nil)
}
