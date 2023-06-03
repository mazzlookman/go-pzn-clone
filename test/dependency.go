package test

import (
	"go-pzn-clone/app"
	"go-pzn-clone/repository"
)

var (
	db       = app.DBConnection()
	userRepo = repository.NewUserRepository(db)
)
