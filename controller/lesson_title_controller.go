package controller

import "github.com/gin-gonic/gin"

type LessonTitleController interface {
	CreateLessonTitle(ctx *gin.Context)
	UpdateLessonTitle(ctx *gin.Context)
	GetLessonTitleByCourseID(ctx *gin.Context)
}
