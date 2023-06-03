package web

type UserRegisterInput struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type UserLoginInput struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID     int
	Name   string
	Email  string
	Avatar string
	Token  string
}
