package service

import "go-pzn-clone/model/web"

type LessonContentService interface {
	Create(input web.LessonContentInput) (web.LessonContentResponse, error)
	Update(lcID int, input web.LessonContentInput) (web.LessonContentResponse, error)
	FindByLessonTitleID(ltID int) ([]web.LessonContentResponse, error)
}
