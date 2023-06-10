package test

import (
	"go-pzn-clone/app"
	"go-pzn-clone/middleware/auth"
	"go-pzn-clone/repository"
	"go-pzn-clone/service"
	"testing"
)

var (
	db          = app.DBConnection()
	jwtAuth     = auth.NewJWTAuth()
	userRepo    = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo, jwtAuth)

	crsRepo    = repository.NewCourseRepository(db)
	crsService = service.NewCourseService(crsRepo)

	ltRepo = repository.NewLessonTitleRepository(db)
)

func logError(err error, t *testing.T) {
	if err != nil {
		t.Log(err.Error())
	}
}
