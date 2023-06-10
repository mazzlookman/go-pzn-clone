package test

import (
	"github.com/stretchr/testify/assert"
	"go-pzn-clone/model/domain"
	"testing"
)

func TestCreateLC(t *testing.T) {
	lessonContent := domain.LessonContent{}
	lessonContent.LessonTitleID = 1
	lessonContent.Content = "contents/golang_dasar_3.mp4"

	content, err := lcRepo.Save(lessonContent)
	logError(err, t)

	assert.Equal(t, "contents/golang_dasar_1.mp4", content.Content)
	t.Log(content)
}

func TestUpdateLC(t *testing.T) {
	findByID, err := lcRepo.FindByID(1)
	logError(err, t)

	findByID.Content = "contents/golang_dasar_1_updated.mp4"

	lessonContent, err := lcRepo.Update(findByID)
	logError(err, t)

	assert.Equal(t, "contents/golang_dasar_1_updated.mp4", lessonContent.Content)
	assert.Equal(t, 1, lessonContent.ID)

	t.Log(lessonContent)
}

func TestFindByLTID(t *testing.T) {
	lessonContents, err := lcRepo.FindByLessonTitleID(1)
	logError(err, t)

	assert.Equal(t, 3, len(lessonContents))
	for _, lessonContent := range lessonContents {
		t.Log(lessonContent)
	}
}
