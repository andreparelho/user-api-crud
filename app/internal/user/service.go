package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/andreparelho/user-api-crud/pkg/dynamo"
	"github.com/andreparelho/user-api-crud/pkg/repository"
	"github.com/gofiber/fiber/v3"
)

type UserService interface {
	CreateUser() fiber.Handler
	GetUser() fiber.Handler
}

func NewUserService(repo repository.UserRepository, logger *slog.Logger) UserService {
	return &user{
		Repository: repo,
		Logger:     logger,
	}
}

func (u *user) CreateUser() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var user User
		if err := json.Unmarshal(ctx.Body(), &user); err != nil {
			u.Logger.ErrorContext(ctx.Context(), "Erro ao fazer o unmarshall do usuario.",
				slog.Any("json", (ctx.Body())),
				slog.Any("error", err),
			)
			return ctx.SendStatus(http.StatusBadRequest)
		}

		u.SetIdToUser(&user)

		dynamoUser := dynamo.User{
			ID:    user.Id,
			Name:  user.Nome,
			Email: user.Email,
		}
		if err := u.Repository.Save(ctx.Context(), dynamoUser); err != nil {
			u.Logger.ErrorContext(ctx.Context(), "Erro ao salvar o usuario no dynamo.",
				slog.Any("error", err),
			)
			return ctx.SendStatus(http.StatusInternalServerError)
		}

		return ctx.SendStatus(http.StatusAccepted)
	}
}

func (u *user) GetUser() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"id": "123",
		})
	}
}
