package test

import (
	"go-pzn-clone/app"
	"go-pzn-clone/repository"
	"go-pzn-clone/service"
)

var (
	db          = app.DBConnection()
	userRepo    = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo)
)
