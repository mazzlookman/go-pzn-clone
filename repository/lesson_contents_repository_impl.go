package repository

import (
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"gorm.io/gorm"
)

type LessonContentRepositoryImpl struct {
	db *gorm.DB
}

func (r *LessonContentRepositoryImpl) FindByID(lcID int) (domain.LessonContent, error) {
	lessonContent := domain.LessonContent{}
	err := r.db.Find(&lessonContent, "id=?", lcID).Error
	helper.PanicIfError(err)

	return lessonContent, nil
}

func (r *LessonContentRepositoryImpl) Save(content domain.LessonContent) (domain.LessonContent, error) {
	err := r.db.Create(&content).Error
	helper.PanicIfError(err)

	return content, nil
}

func (r *LessonContentRepositoryImpl) Update(content domain.LessonContent) (domain.LessonContent, error) {
	err := r.db.Save(&content).Error
	helper.PanicIfError(err)

	return content, nil
}

func (r *LessonContentRepositoryImpl) FindByLessonTitleID(ltID int) ([]domain.LessonContent, error) {
	lessonContents := []domain.LessonContent{}
	err := r.db.Order("in_order asc").Find(&lessonContents, "lesson_title_id=?", ltID).Error
	helper.PanicIfError(err)

	return lessonContents, nil
}

func NewLessonContentRepository(db *gorm.DB) *LessonContentRepositoryImpl {
	return &LessonContentRepositoryImpl{db: db}
}
