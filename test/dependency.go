package test

import (
	"go-pzn-clone/app"
	"go-pzn-clone/middleware/auth"
	"go-pzn-clone/repository"
	"go-pzn-clone/service"
)

var (
	db          = app.DBConnection()
	jwtAuth     = auth.NewJWTAuth()
	userRepo    = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo, jwtAuth)

	courseRepo = repository.NewCourseRepository(db)
)
