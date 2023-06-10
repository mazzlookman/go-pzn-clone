package test

import (
	"github.com/stretchr/testify/assert"
	"go-pzn-clone/model/web"
	"testing"
)

func TestCourseCreate(t *testing.T) {
	input := web.CourseInput{
		AuthorID:    1,
		Title:       "Java Pemula-Mahir",
		Slug:        "java-pemula-mahir",
		Description: "Descripsi Panjang",
		Perks:       "perks1, perks2, perks3",
		Price:       99000,
	}

	create, err := crsService.Create(input)
	if err != nil {
		t.Log(err.Error())
	}

	assert.Equal(t, "Java Pemula-Mahir", create.Title)
	t.Log(create.ID)
}

func TestCourseUpdate(t *testing.T) {
	input := web.CourseInput{
		AuthorID:    1,
		Title:       "Java Pemula-Mahir-Updated",
		Slug:        "java-pemula-mahir",
		Description: "Descripsi Panjang",
		Perks:       "perks1, perks2, perks3",
		Price:       99000,
	}
	update, _ := crsService.Update(6, input)

	assert.Equal(t, "Java Pemula-Mahir-Updated", update.Title)
	assert.Equal(t, 6, update.ID)
}

func TestCourseFindByUserID(t *testing.T) {
	findByUserID, err := crsService.FindByUserID(4)
	logError(err, t)

	for _, courseResponse := range findByUserID {
		t.Log(courseResponse)
	}
}

func TestCourseFindAll(t *testing.T) {
	courseResponses, err := crsService.FindAll()
	logError(err, t)

	for _, courseResponse := range courseResponses {
		t.Log(courseResponse)
	}
}

func TestCourseFindByCategory(t *testing.T) {
	courseResponses, err := crsService.FindByCategory("devops")
	logError(err, t)

	for _, courseResponse := range courseResponses {
		t.Log(courseResponse)
	}
}

func TestCountUserLearned(t *testing.T) {
	countUsersLearned, err := crsService.CountUsersLearned(4)
	logError(err, t)

	t.Log(countUsersLearned)
}

func TestCourseFindBySlug(t *testing.T) {
	courseResponses, err := crsService.FindBySlug("php")
	logError(err, t)

	for _, courseResponse := range courseResponses {
		t.Log(courseResponse)
	}
}
