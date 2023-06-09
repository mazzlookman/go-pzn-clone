package test

import (
	"github.com/stretchr/testify/assert"
	"go-pzn-clone/model/domain"
	"strings"
	"testing"
)

func TestSaveCourse(t *testing.T) {
	cr := domain.Course{
		AuthorID:    1,
		Title:       "Docker Dasar-Mahir",
		Slug:        "docker-dasar-mahir",
		Description: "Deskripsi Panjang",
		Perks:       "perks1, perks2, perks3",
		Price:       99000,
		Banner:      "images/banner/docker.jpg",
	}

	course, err := crsRepo.Save(cr)
	if err != nil {
		t.Log(err.Error())
	}

	assert.Equal(t, cr.Title, course.Title)
	assert.Equal(t, cr.Banner, course.Banner)
	t.Log(course.ID)
}

func TestUpdateCourse(t *testing.T) {
	cr := domain.Course{
		ID:          4,
		AuthorID:    1,
		Title:       "Docker Dasar-Mahir-Updated",
		Slug:        "docker-dasar-mahir",
		Description: "Deskripsi Panjang",
		Perks:       "perks1, perks2, perks3",
		Price:       99000,
		Banner:      "images/banner/docker.jpg",
	}

	course, err := crsRepo.Update(cr)
	if err != nil {
		t.Log(err.Error())
	}

	assert.Equal(t, cr.Title, course.Title)
	assert.Equal(t, cr.Banner, course.Banner)
	assert.Equal(t, cr.ID, course.ID)
	t.Log(course.ID)
}

func TestFindByID(t *testing.T) {
	course, err := crsRepo.FindByID(4)
	if err != nil {
		t.Log(err.Error())
	}

	assert.Equal(t, course.ID, 4)
	assert.Equal(t, course.Title, "Docker Dasar-Mahir-Updated")
}

func TestFindByUserID(t *testing.T) {
	courses, err := crsRepo.FindByUserID(5)
	if err != nil {
		t.Log(err.Error())
	}

	assert.NotNil(t, courses)
	assert.Equal(t, 2, len(courses))
	assert.Equal(t, true, strings.Contains(courses[0].Title, "Golang"))
	assert.Equal(t, true, strings.Contains(courses[1].Title, "JavaScript"))
}

func TestFindAll(t *testing.T) {
	courses, err := crsRepo.FindAll()
	if err != nil {
		t.Log(err.Error())
	}

	assert.NotNil(t, courses)
	assert.Equal(t, 4, len(courses))
}

func TestCountUsersLearned(t *testing.T) {
	usersLearned, err := crsRepo.CountUsersLearned(3)
	if err != nil {
		t.Log(err.Error())
	}

	t.Log(usersLearned)
}

func TestFindByCategory(t *testing.T) {
	courses, err := crsRepo.FindByCategory("devops")
	if err != nil {
		t.Log(err.Error())
	}

	for _, course := range courses {
		t.Log(course)
	}
}
