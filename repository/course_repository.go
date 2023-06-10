package repository

import "go-pzn-clone/model/domain"

type CourseRepository interface {
	Save(courses domain.Course) (domain.Course, error)
	Update(courses domain.Course) (domain.Course, error)
	FindByID(courseID int) (domain.Course, error)
	FindBySlug(slug string) ([]domain.Course, error)
	FindByUserID(userID int) ([]domain.Course, error)
	FindAll() ([]domain.Course, error)
	FindByCategory(category string) ([]domain.Course, error)
	CountUsersLearned(courseID int) (int, error)
}
