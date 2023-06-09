package service

import (
	"errors"
	"go-pzn-clone/formatter"
	"go-pzn-clone/helper"
	"go-pzn-clone/middleware/auth"
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
	"go-pzn-clone/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceImpl struct {
	repository.UserRepository
	auth.JWTAuth
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

	token, err := s.JWTAuth.GenerateToken(findByEmail.ID)
	helper.PanicIfError(err)

	return formatter.ToUserResponse(findByEmail, token), nil
}

func (s *UserServiceImpl) UploadAvatar(userID int, path string) (web.UserResponse, error) {
	findByID, err := s.UserRepository.FindByID(userID)
	helper.PanicIfError(err)

	findByID.Avatar = path
	findByID.UpdatedAt = time.Now()

	update, err := s.UserRepository.Update(findByID)
	helper.PanicIfError(err)

	token, err := s.JWTAuth.GenerateToken(update.ID)
	helper.PanicIfError(err)

	return formatter.ToUserResponse(update, token), nil
}

func (s *UserServiceImpl) RegisterUser(input web.UserRegisterInput) (web.UserResponse, error) {
	user := domain.User{}
	user.Name = input.Name
	user.Email = input.Email
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	user.Role = input.Role
	if input.Role == "" {
		user.Role = "user"
	}

	findByEmail, _ := s.UserRepository.FindByEmail(user.Email)
	if findByEmail.Email == user.Email {
		return web.UserResponse{}, errors.New("Email has been registered")
	}

	save, err := s.UserRepository.Save(user)
	helper.PanicIfError(err)

	token, err := s.JWTAuth.GenerateToken(save.ID)
	helper.PanicIfError(err)

	return formatter.ToUserResponse(save, token), nil
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

	token, err := s.JWTAuth.GenerateToken(update.ID)
	helper.PanicIfError(err)

	return formatter.ToUserResponse(update, token), nil
}

func (s *UserServiceImpl) FindUserByID(userID int) (web.UserResponse, error) {
	findByID, err := s.UserRepository.FindByID(userID)
	helper.PanicIfError(err)

	if findByID.ID == 0 {
		return web.UserResponse{}, errors.New("User not found")
	}

	token, err := s.JWTAuth.GenerateToken(findByID.ID)
	helper.PanicIfError(err)

	return formatter.ToUserResponse(findByID, token), nil
}

func (s *UserServiceImpl) EmailAvailabilityCheck(email web.EmailAvailability) (bool, error) {
	findByEmail, err := s.UserRepository.FindByEmail(email.Email)
	helper.PanicIfError(err)

	if findByEmail.ID == 0 && findByEmail.Email == "" {
		return true, nil
	}

	return false, errors.New("Email is not available")
}

func (s *UserServiceImpl) DeleteUserByID(userID int) (bool, error) {
	deleteByID, err := s.UserRepository.DeleteByID(userID)
	helper.PanicIfError(err)

	if deleteByID == false {
		return false, err
	}

	return true, nil
}

func NewUserService(userRepository repository.UserRepository, jwtAuth auth.JWTAuth) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, JWTAuth: jwtAuth}
}
