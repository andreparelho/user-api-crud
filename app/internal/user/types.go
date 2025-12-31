package user

type User struct {
	Id    string `json:"-"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type user struct{}
