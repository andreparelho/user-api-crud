package main

import (
	"context"
	"log"

	"github.com/andreparelho/user-api-crud/internal/logger"
	"github.com/andreparelho/user-api-crud/internal/server"
	"github.com/andreparelho/user-api-crud/internal/user"
	"github.com/andreparelho/user-api-crud/pkg/config"
	"github.com/andreparelho/user-api-crud/pkg/dynamo"
	"github.com/andreparelho/user-api-crud/pkg/repository"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(" .env não encontrado, usando variáveis do sistema")
		log.Fatal(err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.New(cfg.Env)
	ctx := server.Shutdown()

	dynamoClient, err := dynamo.NewDynamoClient(context.Background(), *cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(dynamoClient, *cfg)
	userService := user.NewUserService(repo, logg)

	server := server.NewServer(userService, logg)

	if err := server.Start(ctx, cfg.Port); err != nil {
		log.Fatal(err)
	}
}
