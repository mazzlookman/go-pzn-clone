package repository

import (
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"gorm.io/gorm"
)

type LessonTitleRepositoryImpl struct {
	db *gorm.DB
}

func (r *LessonTitleRepositoryImpl) Save(title domain.LessonTitle) (domain.LessonTitle, error) {
	err := r.db.Create(&title).Error
	helper.PanicIfError(err)

	return title, nil
}

func (r *LessonTitleRepositoryImpl) Update(title domain.LessonTitle) (domain.LessonTitle, error) {
	err := r.db.Save(&title).Error
	helper.PanicIfError(err)

	return title, nil
}

func (r *LessonTitleRepositoryImpl) FindByID(lessonTitleID int) (domain.LessonTitle, error) {
	lessonTitle := domain.LessonTitle{}
	err := r.db.Find(&lessonTitle, "id=?", lessonTitleID).Error
	helper.PanicIfError(err)

	return lessonTitle, nil
}

func (r *LessonTitleRepositoryImpl) FindByCourseID(courseID int) ([]domain.LessonTitle, error) {
	lessonTitles := []domain.LessonTitle{}
	err := r.db.Order("in_order asc").Find(&lessonTitles, "course_id=?", courseID).Error
	helper.PanicIfError(err)

	return lessonTitles, nil
}

func NewLessonTitleRepository(db *gorm.DB) *LessonTitleRepositoryImpl {
	return &LessonTitleRepositoryImpl{db: db}
}
