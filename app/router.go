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

	//course endpoint dependency
	courseRepository := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepository)
	courseController := controller.NewCourseController(courseService)

	//lesson titles dependency
	lessonTitleRepository := repository.NewLessonTitleRepository(db)
	lessonTitleService := service.NewLessonTitleService(lessonTitleRepository)
	lessonTitleController := controller.NewLessonTitleController(lessonTitleService)

	//router
	router := gin.Default()
	api := router.Group("/api/v1")

	//user endpoints
	api.POST("/email-checkers", userController.EmailAvailabilityCheck)
	api.POST("/users", userController.RegisterUser)
	api.POST("/users/login", userController.LoginUser)
	api.PUT("/users/avatars", middleware.JWTAuthMiddleware(userAuth, userService), userController.UploadAvatar)
	api.DELETE("/users", middleware.JWTAuthMiddleware(userAuth, userService), userController.DeleteCurrentUser)
	api.GET("/users", middleware.JWTAuthMiddleware(userAuth, userService), userController.GetUserDetail)

	//course endpoints
	api.POST("/courses", middleware.JWTAuthMiddleware(userAuth, userService), courseController.CreateCourse)
	api.GET("/courses", courseController.GetCourseByCategory)
	api.PUT("/courses/:course_id", middleware.JWTAuthMiddleware(userAuth, userService), courseController.UpdateCourse)
	api.GET("/courses/slug", courseController.GetCourseBySlug)
	api.GET("/courses/enrolled", middleware.JWTAuthMiddleware(userAuth, userService), courseController.GetCourseByUserID)
	api.GET("/courses/all", courseController.GetAllCourse)
	api.GET("/courses/:course_id", courseController.CountUserLearned)
	api.PUT("/courses/:course_id/banner", middleware.JWTAuthMiddleware(userAuth, userService), courseController.UploadBanner)

	//lesson title endpoints
	api.POST("/courses/lessons", lessonTitleController.CreateLessonTitle)
	api.PUT("/courses/lessons/:lt_id", lessonTitleController.UpdateLessonTitle)
	api.GET("/courses/lessons/:course_id", lessonTitleController.GetLessonTitleByCourseID)

	return router
}
