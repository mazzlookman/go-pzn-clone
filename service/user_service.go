package service

import "go-pzn-clone/model/web"

type UserService interface {
	RegisterUser(input web.UserRegisterInput) (web.UserResponse, error)
	LoginUser(input web.UserLoginInput) (web.UserResponse, error)
	UpdateUser(userID int, input web.UserRegisterInput) (web.UserResponse, error)
	FindUserByID(userID int) (web.UserResponse, error)
	EmailAvailabilityCheck(email web.EmailAvailability) (bool, error)
	DeleteUserByID(userID int) (bool, error)
	UploadAvatar(userID int, path string) (web.UserResponse, error)
}
