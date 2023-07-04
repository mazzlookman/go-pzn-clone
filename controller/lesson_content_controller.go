package controller

import "github.com/gin-gonic/gin"

type LessonContentController interface {
	CreateLessonContent(ctx *gin.Context)
	UpdateLessonContent(ctx *gin.Context)
	GetByLessonTitleID(ctx *gin.Context)
}
