package server

import (
	"time"

	"github.com/andreparelho/user-api-crud/internal/user"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func CreateRouter(userService user.UserService) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	app.Use(recover.New())
	app.Use(requestid.New())

	app.Get("/user", userService.GetUser())
	app.Post("/user", userService.CreateUser())

	return app
}
