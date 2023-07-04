package test

import (
	"github.com/stretchr/testify/assert"
	"go-pzn-clone/model/web"
	"testing"
)

func TestCreateLCService(t *testing.T) {
	lessonContentInput := web.LessonContentInput{
		LessonTitleID: 1,
		InOrder:       1,
		Content:       "resources/contents/exp_1.mp4",
	}

	lessonContentResponse, err := lcService.Create(lessonContentInput)
	logError(err, t)

	t.Log(lessonContentResponse)
}

func TestUpdateLCService(t *testing.T) {
	lessonContentInput := web.LessonContentInput{
		LessonTitleID: 1,
		InOrder:       2,
		Content:       "resources/contents/exp_1.mp4",
	}

	lessonContentResponse, err := lcService.Create(lessonContentInput)
	logError(err, t)

	lessonContentInput.Content = "resources/contents/exp_2.mp4"

	updated, err := lcService.Update(lessonContentResponse.ID, lessonContentInput)
	logError(err, t)

	assert.Equal(t, lessonContentResponse.ID, updated.ID)
	assert.Equal(t, "resources/contents/exp_2.mp4", updated.Content)
}

func TestFindLCByLTID(t *testing.T) {
	lessonContentResponses, err := lcService.FindByLessonTitleID(4)
	logError(err, t)

	for _, response := range lessonContentResponses {
		t.Log(response)
	}
}
