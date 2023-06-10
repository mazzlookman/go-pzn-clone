package controller

import "github.com/gin-gonic/gin"

type CourseController interface {
	CreateCourse(ctx *gin.Context)
	UpdateCourse(ctx *gin.Context)
	GetCourseBySlug(ctx *gin.Context)
	GetCourseByUserID(ctx *gin.Context)
	GetCourseByCategory(ctx *gin.Context)
	GetAllCourse(ctx *gin.Context)
	CountUserLearned(ctx *gin.Context)
	UploadBanner(ctx *gin.Context)
}
