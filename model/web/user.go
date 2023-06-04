package web

type UserRegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string
}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Token  string `json:"token"`
}

type EmailAvailability struct {
	Email string `json:"email"`
}
