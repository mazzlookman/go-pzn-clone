package helper

import (
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
		Token:  "tokentokentoken",
	}
}