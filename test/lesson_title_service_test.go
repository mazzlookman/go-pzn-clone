package test

import (
	"github.com/stretchr/testify/assert"
	"go-pzn-clone/model/web"
	"log"
	"testing"
)

func TestCreateLT(t *testing.T) {
	input := web.LessonTitleInput{
		CourseID: 1,
		InOrder:  3,
		Title:    "Golang Unit Test",
	}

	lessonTitleResponse, err := ltService.Create(input)
	logError(err, t)

	assert.Equal(t, "Golang Unit Test", lessonTitleResponse.Title)
	log.Println(lessonTitleResponse)
}

func TestUpdateLT(t *testing.T) {
	input := web.LessonTitleInput{
		CourseID: 1,
		InOrder:  3,
		Title:    "Golang Unit Test",
	}

	titleResponse, err := ltService.Update(3, input)
	logError(err, t)

	assert.Equal(t, "Golang Unit Test", titleResponse.Title)
	assert.Equal(t, 3, titleResponse.ID)
	log.Println(titleResponse)
}

func TestFindByCourseIDLT(t *testing.T) {
	lessonTitleResponses, err := ltService.FindByCourseID(2)
	logError(err, t)

	//assert.Equal(t, 3, len(lessonTitleResponses))
	for _, lessonTitleResponse := range lessonTitleResponses {
		t.Log(lessonTitleResponse)
	}
}
