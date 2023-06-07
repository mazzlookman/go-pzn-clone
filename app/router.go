package app

import (
	"github.com/gin-gonic/gin"
	"go-pzn-clone/controller"
	"go-pzn-clone/middleware"
	"go-pzn-clone/middleware/auth"
	"go-pzn-clone/repository"
	"go-pzn-clone/service"
)

func RouterInitialized() *gin.Engine {
	db := DBConnection()

	//user endpoint dependency
	userAuth := auth.NewJWTAuth()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, userAuth)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	//user endpoint
	api.POST("/email-checkers", userController.EmailAvailabilityCheck)
	api.POST("/users", userController.RegisterUser)
	api.POST("/users/login", middleware.JWTAuthMiddleware(userAuth, userService), userController.LoginUser)
	api.PUT("/users/avatars", middleware.JWTAuthMiddleware(userAuth, userService), userController.UploadAvatar)
	api.DELETE("/users", middleware.JWTAuthMiddleware(userAuth, userService), userController.DeleteCurrentUser)
	api.GET("/users", middleware.JWTAuthMiddleware(userAuth, userService), userController.GetUserDetail)

	return router
}
