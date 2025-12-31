package user

type UserService interface {
	CreateUser()
	GetUser()
}

func NewUserService() *user {
	return &user{}
}

func (u user) CreateUser() {
}

func (u user) GetUser() {
}
