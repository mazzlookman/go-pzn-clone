package service

import "go-pzn-clone/model/web"

type CourseService interface {
	Create(input web.CourseInput) (web.CourseResponse, error)
	Update(courseID int, input web.CourseInput) (web.CourseResponse, error)
	FindByID(courseID int) (web.CourseResponse, error)
	FindByUserID(userID int) ([]web.CourseResponse, error)
	FindAll() ([]web.CourseResponse, error)
	FindByCategory(category string) ([]web.CourseResponse, error)
	CountUserLearned(courseID int) (int, error)
}
