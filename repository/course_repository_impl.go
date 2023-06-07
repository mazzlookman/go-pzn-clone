package repository

import (
	"go-pzn-clone/model/domain"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db *gorm.DB
}

func (r *CourseRepositoryImpl) Save(course domain.Course) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CourseRepositoryImpl) Update(course domain.Course) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CourseRepositoryImpl) FindByID(courseID int) (domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CourseRepositoryImpl) FindByUserID(userID int) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CourseRepositoryImpl) FindAll() ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CourseRepositoryImpl) FindByCategoryID(categoryID int) ([]domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CourseRepositoryImpl) CountUserLiked(courseID int) (int, error) {
	//TODO implement me
	panic("implement me")
}
