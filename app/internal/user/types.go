package user

import (
	"log/slog"

	"github.com/andreparelho/user-api-crud/pkg/repository"
	"github.com/google/uuid"
)

type User struct {
	Id    string `json:"-"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type user struct {
	Repository repository.UserRepository
	Logger     *slog.Logger
}

func (u user) SetIdToUser(user *User) {
	user.Id = uuid.New().String()
}
