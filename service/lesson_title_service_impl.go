package service

import (
	"go-pzn-clone/formatter"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
	"go-pzn-clone/repository"
	"time"
)

type LessonTitleServiceImpl struct {
	repository.LessonTitleRepository
}

func (s *LessonTitleServiceImpl) Create(input web.LessonTitleInput) (web.LessonTitleResponse, error) {
	lessonTitle := domain.LessonTitle{
		CourseID: input.CourseID,
		InOrder:  input.InOrder,
		Title:    input.Title,
	}
	save, err := s.LessonTitleRepository.Save(lessonTitle)
	helper.PanicIfError(err)

	return formatter.ToLessonTitlesResponse(save), nil
}

func (s *LessonTitleServiceImpl) Update(ltID int, input web.LessonTitleInput) (web.LessonTitleResponse, error) {
	findByID, err := s.LessonTitleRepository.FindByID(ltID)
	helper.PanicIfError(err)
	findByID.CourseID = input.CourseID
	findByID.InOrder = input.InOrder
	findByID.Title = input.Title
	findByID.UpdatedAt = time.Now()

	lessonTitle, err := s.LessonTitleRepository.Update(findByID)
	helper.PanicIfError(err)

	return formatter.ToLessonTitlesResponse(lessonTitle), nil
}

func (s *LessonTitleServiceImpl) FindByCourseID(courseID int) ([]web.LessonTitleResponse, error) {
	lessonTitles, err := s.LessonTitleRepository.FindByCourseID(courseID)
	helper.PanicIfError(err)

	return formatter.ToLessonTitlesResponses(lessonTitles), nil
}

func NewLessonTitleService(lessonTitleRepository repository.LessonTitleRepository) *LessonTitleServiceImpl {
	return &LessonTitleServiceImpl{LessonTitleRepository: lessonTitleRepository}
}
