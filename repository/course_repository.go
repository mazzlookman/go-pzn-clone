package repository

import "go-pzn-clone/model/domain"

type CourseRepository interface {
	Save(courses domain.Course) (domain.Course, error)
	Update(courses domain.Course) (domain.Course, error)
	FindByID(courseID int) (domain.Course, error)
	FindByUserID(userID int) ([]domain.Course, error)
	FindAll() ([]domain.Course, error)
	FindByCategoryID(categoryID int) ([]domain.Course, error)
	CountUserLiked(courseID int) (int, error)
}
