package service

import (
	"go-pzn-clone/formatter"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
	"go-pzn-clone/repository"
	"time"
)

type LessonContentServiceImpl struct {
	repository.LessonContentRepository
}

func (s *LessonContentServiceImpl) Create(input web.LessonContentInput) (web.LessonContentResponse, error) {
	lc := domain.LessonContent{}
	lc.LessonTitleID = input.LessonTitleID
	lc.InOrder = input.InOrder
	lc.Content = input.Content

	lessonContent, err := s.LessonContentRepository.Save(lc)
	helper.PanicIfError(err)

	return formatter.ToLessonContentResponse(lessonContent), nil
}

func (s *LessonContentServiceImpl) Update(lcID int, input web.LessonContentInput) (web.LessonContentResponse, error) {
	findByID, err := s.LessonContentRepository.FindByID(lcID)
	helper.PanicIfError(err)
	findByID.LessonTitleID = input.LessonTitleID
	findByID.InOrder = input.InOrder
	findByID.Content = input.Content
	findByID.UpdatedAt = time.Now()

	lessonContent, err := s.LessonContentRepository.Update(findByID)
	helper.PanicIfError(err)

	return formatter.ToLessonContentResponse(lessonContent), nil
}

func (s *LessonContentServiceImpl) FindByLessonTitleID(ltID int) ([]web.LessonContentResponse, error) {
	lessonContents, err := s.LessonContentRepository.FindByLessonTitleID(ltID)
	helper.PanicIfError(err)

	return formatter.ToLessonContentResponses(lessonContents), nil
}

func NewLessonContentService(lessonContentRepository repository.LessonContentRepository) *LessonContentServiceImpl {
	return &LessonContentServiceImpl{LessonContentRepository: lessonContentRepository}
}
