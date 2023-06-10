package test

import (
	"github.com/stretchr/testify/assert"
	"go-pzn-clone/model/domain"
	"log"
	"testing"
)

func TestCreateLessonTitle(t *testing.T) {
	lt := domain.LessonTitle{
		CourseID: 1,
		InOrder:  2,
		Title:    "Golang Modules",
	}

	lessonTitle, err := ltRepo.Save(lt)
	logError(err, t)

	assert.Equal(t, "Golang Modules", lessonTitle.Title)
	assert.Equal(t, 2, lessonTitle.InOrder)
	log.Println(lessonTitle)
}

func TestUpdateLessonTitle(t *testing.T) {
	findByID, _ := ltRepo.FindByID(2)
	findByID.InOrder = 2
	findByID.Title = "Golang Modules"

	lessonTitle, err := ltRepo.Update(findByID)
	logError(err, t)

	assert.Equal(t, "Golang Modules", lessonTitle.Title)
	assert.Equal(t, 2, lessonTitle.InOrder)
	assert.Equal(t, 2, lessonTitle.ID)
	log.Println(lessonTitle)
}

func TestFindByCourseID(t *testing.T) {
	lessonTitles, err := ltRepo.FindByCourseID(1)
	logError(err, t)

	assert.Equal(t, 2, len(lessonTitles))
	assert.Equal(t, "Golang Modules", lessonTitles[1].Title)
	assert.Equal(t, "Golang Dasar", lessonTitles[0].Title)

	for _, lessonTitle := range lessonTitles {
		log.Println(lessonTitle)
	}

}
