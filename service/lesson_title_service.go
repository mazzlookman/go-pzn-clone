package service

import "go-pzn-clone/model/web"

type LessonTitleService interface {
	Create(input web.LessonTitleInput) (web.LessonTitleResponse, error)
	Update(ltID int, input web.LessonTitleInput) (web.LessonTitleResponse, error)
	FindByCourseID(courseID int) ([]web.LessonTitleResponse, error)
}
