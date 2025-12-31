package server

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App    *fiber.App
	Logger *slog.Logger
}
