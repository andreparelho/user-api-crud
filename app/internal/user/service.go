package user

import "github.com/gofiber/fiber/v3"

type UserService interface {
	CreateUser() fiber.Handler
	GetUser() fiber.Handler
}

func NewUserService() UserService {
	return &user{}
}

func (u *user) CreateUser() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"ok": true,
		})
	}
}

func (u *user) GetUser() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"id": "123",
		})
	}
}
