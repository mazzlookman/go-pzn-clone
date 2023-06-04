package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pzn-clone/formatter"
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
	"go-pzn-clone/service"
	"net/http"
)

type UserControllerImpl struct {
	service.UserService
}

func (c *UserControllerImpl) EmailAvailabilityCheck(ctx *gin.Context) {
	inputJSON := web.EmailAvailability{}
	err := ctx.ShouldBindJSON(&inputJSON)
	if err != nil {
		apiResponse := formatter.APIResponse("Email checking failed", http.StatusBadRequest, "BAD REQUEST", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	emailAvailabilityCheck, err := c.UserService.EmailAvailabilityCheck(inputJSON)
	if err != nil {
		apiResponse := formatter.APIResponse("Email checking failed", http.StatusInternalServerError, "INTERNAL SERVER ERROR", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	var metaMsg string
	if emailAvailabilityCheck {
		metaMsg = "Email is available"
	} else {
		metaMsg = "Email has been registered"
	}

	ctx.JSON(200, formatter.APIResponse(
		metaMsg,
		200,
		"success",
		gin.H{"is_available": emailAvailabilityCheck},
	))
}

func (c *UserControllerImpl) RegisterUser(ctx *gin.Context) {
	inputJSON := web.UserRegisterInput{}
	err := ctx.ShouldBindJSON(&inputJSON)
	if err != nil {
		apiResponse := formatter.APIResponse("Register is failed", 400, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	registerUser, err := c.UserService.RegisterUser(inputJSON)
	if err != nil {
		apiResponse := formatter.APIResponse("Register is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Register is successfully", 200, "success", registerUser)
	ctx.JSON(200, apiResponse)
}

func (c *UserControllerImpl) LoginUser(ctx *gin.Context) {
	inputJSON := web.UserLoginInput{}
	err := ctx.ShouldBindJSON(&inputJSON)
	if err != nil {
		apiResponse := formatter.APIResponse("Login is failed", 400, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, apiResponse)
	}

	loginUser, err := c.UserService.LoginUser(inputJSON)
	if err != nil {
		apiResponse := formatter.APIResponse("Login is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
	}

	apiResponse := formatter.APIResponse("Register is successfully", 200, "success", loginUser)
	ctx.JSON(200, apiResponse)
}

func (c *UserControllerImpl) UploadAvatar(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("avatar")
	if err != nil {
		apiResponse := formatter.APIResponse("Failed to upload avatar", http.StatusBadRequest, "BAD REQUEST", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	userID := 1
	path := fmt.Sprintf("images/avatar/%d-%s", userID, fileHeader.Filename)

	_ = ctx.SaveUploadedFile(fileHeader, path)

	_, err = c.UserService.UploadAvatar(userID, path)
	if err != nil {
		apiResponse := formatter.APIResponse("Failed to upload avatar", http.StatusInternalServerError, "INTERNAL SERVER ERROR", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	ctx.JSON(200, formatter.APIResponse(
		"Avatar is successfully uploaded",
		200,
		"success",
		gin.H{"is_uploaded": true},
	))
}

func (c *UserControllerImpl) GetUserDetail(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(domain.User)
	apiResponse := formatter.APIResponse("Detail of current user", 200, "success", user)
	ctx.JSON(200, apiResponse)
}

func (c *UserControllerImpl) DeleteCurrentUser(ctx *gin.Context) {
	userID := ctx.MustGet("currentUser").(domain.User).ID
	deleteUserByID, err := c.UserService.DeleteUserByID(userID)
	if err != nil {
		apiResponse := formatter.APIResponse("Failed to delete user", http.StatusInternalServerError, "INTERNAL SERVER ERROR", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("User is successfully deleted", 200, "success", gin.H{"is_deleted": deleteUserByID})
	ctx.JSON(200, apiResponse)
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}
