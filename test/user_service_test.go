package test

import (
	"github.com/stretchr/testify/assert"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/web"
	"log"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	input := web.UserRegisterInput{
		Name:     "Ucup",
		Email:    "ucup@test.com",
		Password: "123",
	}

	userResponse, err := userService.RegisterUser(input)
	helper.PanicIfError(err)

	findUserByID, _ := userService.FindUserByID(userResponse.ID)

	assert.Equal(t, "Ucup", userResponse.Name)
	assert.Equal(t, userResponse.ID, findUserByID.ID)
}

func TestLoginUser(t *testing.T) {
	input := web.UserLoginInput{
		Email:    "ucup@test.com",
		Password: "123",
	}

	loginUser, err := userService.LoginUser(input)
	helper.PanicIfError(err)

	assert.Equal(t, loginUser.Email, "ucup@test.com")
	assert.NotEqual(t, 0, loginUser.ID)
}

func TestLoginUserFailed(t *testing.T) {
	input := web.UserLoginInput{
		Email:    "ucup@test.comm",
		Password: "123s",
	}

	loginUser, err := userService.LoginUser(input)

	assert.Equal(t, 0, loginUser.ID)
	assert.Equal(t, "Login failed", err.Error())

	log.Println(err.Error())
}

func TestUploadAvatar(t *testing.T) {
	uploadAvatar, err := userService.UploadAvatar(4, "image/ucup.jpg")

	assert.Nil(t, err)
	assert.NotNil(t, uploadAvatar.ID)
	assert.Equal(t, uploadAvatar.Avatar, "image/ucup.jpg")
}

func TestFindUserByID(t *testing.T) {
	user, err := userService.FindUserByID(4)

	//assert.Equal(t, "User not found", err.Error())
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
}

func TestFindUserByEmail(t *testing.T) {
	user, err := userService.FindUserByEmail("ucup@test.com")

	//assert.Equal(t, "User not found", err.Error())
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
}

func TestDeleteUserByID(t *testing.T) {
	findUserByID, _ := userService.FindUserByID(5)
	deleteUserByID := userService.DeleteUserByID(findUserByID.ID)
	assert.Equal(t, "User was deleted", deleteUserByID)

	userDeleted, _ := userService.FindUserByID(5)
	assert.Equal(t, 0, userDeleted.ID)
}
