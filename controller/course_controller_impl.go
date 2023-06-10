package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pzn-clone/formatter"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/web"
	"go-pzn-clone/service"
	"net/http"
)

type CourseControllerImpl struct {
	service.CourseService
}

func (c *CourseControllerImpl) UploadBanner(ctx *gin.Context) {
	cid := web.CourseIDFromURI{}
	err := ctx.ShouldBindUri(&cid)
	if err != nil {
		apiResponse := formatter.APIResponse("Upload course banner is failed", 400, "error", helper.ValidationError(err))
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	fileHeader, err := ctx.FormFile("banner")
	helper.PanicIfError(err)

	path := fmt.Sprintf("images/banner/%d-%s", cid.ID, fileHeader.Filename)
	err = ctx.SaveUploadedFile(fileHeader, path)
	helper.PanicIfError(err)

	courseResponse, err := c.CourseService.UploadBanner(cid.ID, path)
	if err != nil {
		apiResponse := formatter.APIResponse("Upload course banner is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Course banner successfully uploaded", 200, "success", courseResponse)
	ctx.JSON(200, apiResponse)
}

func (c *CourseControllerImpl) CreateCourse(ctx *gin.Context) {
	input := web.CourseInput{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		apiResponse := formatter.APIResponse("Create course is failed", 400, "error", helper.ValidationError(err))
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	courseResponse, err := c.CourseService.Create(input)
	if err != nil {
		apiResponse := formatter.APIResponse("Create course is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Course successfully created", 200, "success", courseResponse)
	ctx.JSON(200, apiResponse)
}

func (c *CourseControllerImpl) UpdateCourse(ctx *gin.Context) {
	cid := web.CourseIDFromURI{}
	err := ctx.ShouldBindUri(&cid)

	input := web.CourseInput{}
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		apiResponse := formatter.APIResponse("Update course is failed", 400, "error", helper.ValidationError(err))
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	courseResponse, err := c.CourseService.Update(cid.ID, input)
	if err != nil {
		apiResponse := formatter.APIResponse("Update course is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Course successfully updated", 200, "success", courseResponse)
	ctx.JSON(200, apiResponse)
}

func (c *CourseControllerImpl) GetCourseBySlug(ctx *gin.Context) {
	slug := web.CourseSlug{}
	err := ctx.ShouldBindJSON(&slug)
	if err != nil {
		apiResponse := formatter.APIResponse("Update course is failed", 400, "error", helper.ValidationError(err))
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	courseResponses, err := c.CourseService.FindBySlug(slug.Slug)
	if err != nil {
		apiResponse := formatter.APIResponse("Update course is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Course successfully updated", 200, "success", courseResponses)
	ctx.JSON(200, apiResponse)
}

func (c *CourseControllerImpl) GetCourseByUserID(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(web.UserResponse)

	courseResponses, err := c.CourseService.FindByUserID(user.ID)
	if err != nil {
		apiResponse := formatter.APIResponse("Get course is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("List of course of user "+user.Name, 200, "success", courseResponses)
	ctx.JSON(200, apiResponse)
}

func (c *CourseControllerImpl) GetCourseByCategory(ctx *gin.Context) {
	queryCategory := ctx.Query("category")

	courseResponses, err := c.CourseService.FindByCategory(queryCategory)
	if err != nil {
		apiResponse := formatter.APIResponse("Get course is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("List of course category "+queryCategory, 200, "success", courseResponses)
	ctx.JSON(200, apiResponse)
}

func (c *CourseControllerImpl) GetAllCourse(ctx *gin.Context) {
	courseResponses, err := c.CourseService.FindAll()
	if err != nil {
		apiResponse := formatter.APIResponse("Get course is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("List all course", 200, "success", courseResponses)
	ctx.JSON(200, apiResponse)
}

func (c *CourseControllerImpl) CountUserLearned(ctx *gin.Context) {
	cid := web.CourseIDFromURI{}
	err := ctx.ShouldBindUri(&cid)
	helper.PanicIfError(err)

	countUserLearned, err := c.CourseService.CountUserLearned(cid.ID)
	helper.PanicIfError(err)

	ctx.JSON(200, gin.H{"users_learned": countUserLearned})
}

func NewCourseController(courseService service.CourseService) *CourseControllerImpl {
	return &CourseControllerImpl{CourseService: courseService}
}
