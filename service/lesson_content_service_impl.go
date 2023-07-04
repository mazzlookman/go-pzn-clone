package service

import (
	"go-pzn-clone/formatter"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
	"go-pzn-clone/repository"
	"log"
	"time"
)

type LessonContentServiceImpl struct {
	repository.LessonContentRepository
}

func (s *LessonContentServiceImpl) FindByID(lcID int) (domain.LessonContent, error) {
	lessonContent, err := s.LessonContentRepository.FindByID(lcID)
	helper.PanicIfError(err)

	return lessonContent, nil
}

func (s *LessonContentServiceImpl) Create(input web.LessonContentInput) (web.LessonContentResponse, error) {
	lc := domain.LessonContent{}
	lc.LessonTitleID = input.LessonTitleID
	lc.InOrder = input.InOrder
	lc.Content = input.Content
	lc.Duration = input.Duration

	lessonContent, err := s.LessonContentRepository.Save(lc)
	helper.PanicIfError(err)

	return formatter.ToLessonContentResponse(lessonContent), nil
}

func (s *LessonContentServiceImpl) Update(lciD int, input web.LessonContentInput) (web.LessonContentResponse, error) {
	content, err2 := s.LessonContentRepository.FindByID(lciD)
	helper.PanicIfError(err2)
	oldContent := content.Content

	content.LessonTitleID = input.LessonTitleID
	if input.InOrder != 0 {
		content.InOrder = input.InOrder
	}

	if input.Content != "" {
		content.Content = input.Content
	}

	content.Duration = input.Duration
	content.UpdatedAt = time.Now()

	lessonContent, err := s.LessonContentRepository.Update(content)
	helper.PanicIfError(err)

	if oldContent != input.Content {
		deleteLessonContent := helper.DeleteLessonContent(oldContent)
		log.Println(deleteLessonContent)
	}

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
