package service

import (
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
)

type LessonContentService interface {
	Create(input web.LessonContentInput) (web.LessonContentResponse, error)
	Update(lciD int, input web.LessonContentInput) (web.LessonContentResponse, error)
	FindByID(lcID int) (domain.LessonContent, error)
	FindByLessonTitleID(ltID int) ([]web.LessonContentResponse, error)
}
