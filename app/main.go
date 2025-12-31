package main

import (
	"log"

	"github.com/andreparelho/user-api-crud/app/config"
	"github.com/andreparelho/user-api-crud/app/internal/logger"
	"github.com/andreparelho/user-api-crud/app/internal/server"
)

func main() {
	cfg := config.Load()
	logg := logger.New(cfg.Env)
	ctx := server.Shutdown()

	app := server.CreateRouter()
	server := server.NewServer(app, logg)

	if err := server.Start(ctx, cfg.Port); err != nil {
		log.Fatal(err)
	}
}
