package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	UploadAvatar(ctx *gin.Context)
	GetUserDetail(ctx *gin.Context)
	EmailAvailabilityCheck(ctx *gin.Context)
	DeleteCurrentUser(ctx *gin.Context)
}
