package service

import (
	"errors"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
	"go-pzn-clone/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (s *UserServiceImpl) LoginUser(input web.UserLoginInput) (web.UserResponse, error) {
	findByEmail, err := s.UserRepository.FindByEmail(input.Email)
	helper.PanicIfError(err)

	if findByEmail.ID == 0 {
		return web.UserResponse{}, errors.New("Login failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.Password), []byte(input.Password))
	if err != nil {
		return web.UserResponse{}, errors.New("Login failed")
	}

	return helper.ToUserResponse(findByEmail), nil
}

func (s *UserServiceImpl) UploadAvatar(userID int, path string) (web.UserResponse, error) {
	findByID, err := s.UserRepository.FindByID(userID)
	helper.PanicIfError(err)

	findByID.Avatar = path
	findByID.UpdatedAt = time.Now()

	update, err := s.UserRepository.Update(findByID)
	helper.PanicIfError(err)

	return helper.ToUserResponse(update), nil
}

func (s *UserServiceImpl) RegisterUser(input web.UserRegisterInput) (web.UserResponse, error) {
	user := domain.User{}
	user.Name = input.Name
	user.Email = input.Email
	password, err2 := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err2)
	user.Password = string(password)
	if input.Role == "" {
		user.Role = "user"
	}

	save, err := s.UserRepository.Save(user)
	helper.PanicIfError(err)

	return helper.ToUserResponse(save), nil
}

func (s *UserServiceImpl) UpdateUser(userID int, input web.UserRegisterInput) (web.UserResponse, error) {
	findByID, err := s.UserRepository.FindByID(userID)
	helper.PanicIfError(err)

	findByID.Name = input.Name
	findByID.Email = input.Email
	findByID.Password = input.Password
	findByID.UpdatedAt = time.Now()

	update, err := s.UserRepository.Update(findByID)
	helper.PanicIfError(err)

	return helper.ToUserResponse(update), nil
}

func (s *UserServiceImpl) FindUserByID(userID int) (web.UserResponse, error) {
	findByID, err := s.UserRepository.FindByID(userID)
	helper.PanicIfError(err)

	if findByID.ID == 0 {
		return web.UserResponse{}, errors.New("User not found")
	}

	return helper.ToUserResponse(findByID), nil
}

func (s *UserServiceImpl) FindUserByEmail(email string) (web.UserResponse, error) {
	findByEmail, err := s.UserRepository.FindByEmail(email)
	helper.PanicIfError(err)

	if findByEmail.ID == 0 {
		return web.UserResponse{}, errors.New("User not found")
	}

	return helper.ToUserResponse(findByEmail), nil
}

func (s *UserServiceImpl) DeleteUserByID(userID int) string {
	err := s.UserRepository.DeleteByID(userID)
	helper.PanicIfError(err)

	return "User was deleted"
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository}
}
