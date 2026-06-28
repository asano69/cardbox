package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "my-app/migrations"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// auto-generate a migration file whenever collections change via the Dashboard
		// (enabled only in "go run" dev mode, not in the compiled binary)
		Automigrate: true,
	})

	// Custom routes and hooks go here
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// se.Router.GET("/api/custom", handleCustom)
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
