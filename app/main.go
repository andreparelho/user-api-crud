package main

import (
	"log"

	"github.com/andreparelho/user-api-crud/app/config"
	"github.com/andreparelho/user-api-crud/app/internal/logger"
	"github.com/andreparelho/user-api-crud/app/internal/server"
	"github.com/andreparelho/user-api-crud/app/internal/user"
)

func main() {
	cfg := config.Load()
	logg := logger.New(cfg.Env)
	ctx := server.Shutdown()

	userService := user.NewUserService()

	app := server.CreateRouter(userService)
	server := server.NewServer(app, logg)

	if err := server.Start(ctx, cfg.Port); err != nil {
		log.Fatal(err)
	}
}
