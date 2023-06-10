package repository

import "go-pzn-clone/model/domain"

type LessonTitleRepository interface {
	Save(title domain.LessonTitle) (domain.LessonTitle, error)
	Update(title domain.LessonTitle) (domain.LessonTitle, error)
	FindByID(lessonTitleID int) (domain.LessonTitle, error)
	FindByCourseID(courseID int) ([]domain.LessonTitle, error)
}
