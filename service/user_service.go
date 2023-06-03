package service

import "go-pzn-clone/model/web"

type UserService interface {
	RegisterUser(input web.UserRegisterInput) (web.UserResponse, error)
	LoginUser(input web.UserLoginInput) (web.UserResponse, error)
	UpdateUser(userID int, input web.UserRegisterInput) (web.UserResponse, error)
	FindUserByID(userID int) (web.UserResponse, error)
	FindUserByEmail(email string) (web.UserResponse, error)
	DeleteUserByID(userID int)
	UploadAvatar(userID int, path string) (web.UserResponse, error)
}
