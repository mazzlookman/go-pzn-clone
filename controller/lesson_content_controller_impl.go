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

type LessonContentControllerImpl struct {
	service.LessonContentService
}

func (c *LessonContentControllerImpl) CreateLessonContent(ctx *gin.Context) {
	input := web.LessonContentInput{}
	err := ctx.ShouldBind(&input)
	helper.PanicIfError(err)

	fileHeader, err := ctx.FormFile("content")
	helper.PanicIfError(err)

	path := fmt.Sprintf("resources/contents/%s", fileHeader.Filename)
	err = ctx.SaveUploadedFile(fileHeader, path)
	helper.PanicIfError(err)

	input.Content = path
	input.Duration = helper.GetLessonContentVideoDuration(path)

	lessonContentResponse, err := c.LessonContentService.Create(input)
	if err != nil {
		apiResponse := formatter.APIResponse("Create lesson content is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Lesson content successfully created", 200, "success", lessonContentResponse)
	ctx.JSON(200, apiResponse)
}

func (c *LessonContentControllerImpl) UpdateLessonContent(ctx *gin.Context) {
	lcID := web.LessonContentIDFromURI{}
	err := ctx.ShouldBindUri(&lcID)
	helper.PanicIfError(err)

	input := web.LessonContentInput{}
	err = ctx.ShouldBind(&input)
	helper.PanicIfError(err)

	fileHeader, err := ctx.FormFile("content")
	helper.PanicIfError(err)

	path := fmt.Sprintf("resources/contents/%d-%s", lcID.ID, fileHeader.Filename)
	err = ctx.SaveUploadedFile(fileHeader, path)
	helper.PanicIfError(err)

	input.Content = path
	input.Duration = helper.GetLessonContentVideoDuration(path)

	lessonContentResponse, err := c.LessonContentService.Update(lcID.ID, input)
	if err != nil {
		apiResponse := formatter.APIResponse("Update lesson content is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Lesson content successfully updated", 200, "success", lessonContentResponse)
	ctx.JSON(200, apiResponse)
}

func (c *LessonContentControllerImpl) GetByLessonTitleID(ctx *gin.Context) {
	input := web.LessonTitleIDFromURI{}
	err := ctx.ShouldBindUri(&input)
	helper.PanicIfError(err)

	lessonContentResponses, err := c.LessonContentService.FindByLessonTitleID(input.ID)
	if err != nil {
		apiResponse := formatter.APIResponse("Get lesson content is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Get lesson content successfully created", 200, "success", lessonContentResponses)
	ctx.JSON(200, apiResponse)
}

func NewLessonContentController(lessonContentService service.LessonContentService) *LessonContentControllerImpl {
	return &LessonContentControllerImpl{LessonContentService: lessonContentService}
}
