package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-clone/formatter"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/web"
	"go-pzn-clone/service"
	"net/http"
)

type LessonTitleControllerImpl struct {
	service.LessonTitleService
}

func NewLessonTitleController(lessonTitleService service.LessonTitleService) *LessonTitleControllerImpl {
	return &LessonTitleControllerImpl{LessonTitleService: lessonTitleService}
}

func (c *LessonTitleControllerImpl) CreateLessonTitle(ctx *gin.Context) {
	lessonTitleInput := web.LessonTitleInput{}
	err := ctx.ShouldBindJSON(&lessonTitleInput)
	if err != nil {
		apiResponse := formatter.APIResponse("Create lesson title is failed", 400, "error", helper.ValidationError(err))
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	lessonTitleResponse, err := c.LessonTitleService.Create(lessonTitleInput)
	if err != nil {
		apiResponse := formatter.APIResponse("Create lesson title is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Lesson title successfully created", 200, "success", lessonTitleResponse)
	ctx.JSON(200, apiResponse)
}

func (c *LessonTitleControllerImpl) UpdateLessonTitle(ctx *gin.Context) {
	ltid := web.LessonTitleIDFromURI{}
	err2 := ctx.ShouldBindUri(&ltid)
	helper.PanicIfError(err2)

	lessonTitleInput := web.LessonTitleInput{}
	err := ctx.ShouldBindJSON(&lessonTitleInput)
	if err != nil {
		apiResponse := formatter.APIResponse("Update lesson title is failed", 400, "error", helper.ValidationError(err))
		ctx.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	lessonTitleResponse, err := c.LessonTitleService.Update(ltid.ID, lessonTitleInput)
	if err != nil {
		apiResponse := formatter.APIResponse("Update lesson title is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("Lesson title successfully updated", 200, "success", lessonTitleResponse)
	ctx.JSON(200, apiResponse)
}

func (c *LessonTitleControllerImpl) GetLessonTitleByCourseID(ctx *gin.Context) {
	cid := web.CourseIDFromURI{}
	err := ctx.ShouldBindUri(&cid)
	helper.PanicIfError(err)

	lessonTitleResponses, err := c.LessonTitleService.FindByCourseID(cid.ID)
	if err != nil {
		apiResponse := formatter.APIResponse("Get lesson title is failed", http.StatusInternalServerError, "error", gin.H{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := formatter.APIResponse("List of lesson title by courseID", 200, "success", lessonTitleResponses)
	ctx.JSON(200, apiResponse)
}
