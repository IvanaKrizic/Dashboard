package main

import (
	"github.com/IvanaKrizic/Dashboard/src/config"
	"github.com/IvanaKrizic/Dashboard/src/routes"
)

func main() {
	config.InitDb()
	config.Migrate()

	router := routes.SetupRouter()
	router.Run()
}
