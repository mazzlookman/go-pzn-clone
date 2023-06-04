package app

import (
	"github.com/gin-gonic/gin"
	"go-pzn-clone/controller"
	"go-pzn-clone/repository"
	"go-pzn-clone/service"
)

func RouterInitialized() *gin.Engine {
	db := DBConnection()

	//user endpoint dependency
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	//user endpoint
	api.POST("/email-checkers", userController.EmailAvailabilityCheck)
	api.POST("/users/register", userController.RegisterUser)
	api.POST("/users/login", userController.LoginUser)
	api.PUT("/users/avatars", userController.UploadAvatar)
	api.DELETE("/users/delete", userController.DeleteCurrentUser)
	api.GET("/users/detail", userController.GetUserDetail)

	return router
}
