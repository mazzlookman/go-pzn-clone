package repository

import (
	"context"
	"database/sql"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db *gorm.DB
}

func (r *CourseRepositoryImpl) dbManual() (*sql.DB, context.Context) {
	db, err := r.db.DB()
	helper.PanicIfError(err)

	return db, context.Background()
}

func (r *CourseRepositoryImpl) Save(course domain.Course) (domain.Course, error) {
	err := r.db.Create(&course).Error
	helper.PanicIfError(err)

	return course, nil
}

func (r *CourseRepositoryImpl) Update(course domain.Course) (domain.Course, error) {
	err := r.db.Save(&course).Error
	helper.PanicIfError(err)

	return course, nil
}

func (r *CourseRepositoryImpl) FindByID(courseID int) (domain.Course, error) {
	course := domain.Course{}
	err := r.db.Find(&course, "id=?", courseID).Error
	helper.PanicIfError(err)

	return course, nil
}

func (r *CourseRepositoryImpl) FindByUserID(userID int) ([]domain.Course, error) {
	var courses []domain.Course
	dbManual, ctx := r.dbManual()

	q := "select courses.id, courses.author_id, courses.title, courses.slug, courses.description, courses.perks, courses.price, courses.banner from courses join users_courses on (users_courses.course_id = courses.id) join users on (users_courses.user_id = users.id) where users.id=?"
	rows, err := dbManual.QueryContext(ctx, q, userID)
	defer rows.Close()
	if err != nil {
		return courses, err
	}

	crs := domain.Course{}
	for rows.Next() {
		err := rows.Scan(
			&crs.ID, &crs.AuthorID, &crs.Title, &crs.Slug, &crs.Description, &crs.Perks, &crs.Price, &crs.Banner,
		)
		if err != nil {
			return courses, err
		}
		courses = append(courses, crs)
	}

	return courses, nil
}

func (r *CourseRepositoryImpl) FindAll() ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Order("id desc").Find(&courses).Error
	helper.PanicIfError(err)

	return courses, nil
}

func (r *CourseRepositoryImpl) FindByCategory(category string) ([]domain.Course, error) {
	courses := []domain.Course{}
	query := "select c.id, c.author_id, c.title, c.slug, c.description,c.perks, c.price, c.banner from courses c join categories_courses cc on cc.course_id = c.id join categories ct on cc.category_id = ct.id where ct.name = ?"
	dbManual, ctx := r.dbManual()

	rows, err := dbManual.QueryContext(ctx, query, category)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		c := domain.Course{}
		err := rows.Scan(&c.ID, &c.AuthorID, &c.Title, &c.Slug, &c.Description, &c.Perks, &c.Price, &c.Banner)
		helper.PanicIfError(err)

		courses = append(courses, c)
	}

	return courses, nil
}

func (r *CourseRepositoryImpl) CountUsersLearned(courseID int) (int, error) {
	query := "select count(u.id) from courses as c join users_courses as uc on uc.course_id = c.id join users as u on uc.user_id = u.id where c.id = ?"
	dbManual, ctx := r.dbManual()

	rows, err := dbManual.QueryContext(ctx, query, courseID)
	helper.PanicIfError(err)
	defer rows.Close()

	var total int
	if rows.Next() {
		err := rows.Scan(&total)
		helper.PanicIfError(err)
	}

	return total, nil
}

func NewCourseRepository(db *gorm.DB) *CourseRepositoryImpl {
	return &CourseRepositoryImpl{db: db}
}
