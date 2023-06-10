package repository

import "go-pzn-clone/model/domain"

type LessonContentRepository interface {
	Save(content domain.LessonContent) (domain.LessonContent, error)
	Update(content domain.LessonContent) (domain.LessonContent, error)
	FindByID(lcID int) (domain.LessonContent, error)
	FindByLessonTitleID(ltID int) ([]domain.LessonContent, error)
}
