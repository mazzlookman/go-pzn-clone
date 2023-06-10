package service

import (
	"go-pzn-clone/formatter"
	"go-pzn-clone/helper"
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
	"go-pzn-clone/repository"
	"time"
)

type CourseServiceImpl struct {
	repository.CourseRepository
}

func (s *CourseServiceImpl) UploadBanner(courseID int, path string) (web.CourseResponse, error) {
	findByID, err := s.CourseRepository.FindByID(courseID)
	helper.PanicIfError(err)

	findByID.Banner = path
	findByID.UpdatedAt = time.Now()

	course, err := s.CourseRepository.Update(findByID)
	helper.PanicIfError(err)

	return formatter.ToCourseResponse(course), nil
}

func (s *CourseServiceImpl) FindBySlug(slug string) ([]web.CourseResponse, error) {
	courses, err := s.CourseRepository.FindBySlug(slug)
	helper.PanicIfError(err)

	return formatter.ToCourseResponses(courses), nil
}

func NewCourseService(courseRepository repository.CourseRepository) *CourseServiceImpl {
	return &CourseServiceImpl{CourseRepository: courseRepository}
}

func (s *CourseServiceImpl) Create(input web.CourseInput) (web.CourseResponse, error) {
	course := domain.Course{
		AuthorID:    input.AuthorID,
		Title:       input.Title,
		Slug:        input.Slug,
		Description: input.Description,
		Perks:       input.Perks,
		Price:       input.Price,
	}

	save, err := s.CourseRepository.Save(course)
	helper.PanicIfError(err)

	return formatter.ToCourseResponse(save), nil
}

func (s *CourseServiceImpl) Update(courseID int, input web.CourseInput) (web.CourseResponse, error) {
	findByID, err := s.CourseRepository.FindByID(courseID)
	helper.PanicIfError(err)

	findByID.AuthorID = input.AuthorID
	findByID.Title = input.Title
	findByID.Slug = input.Slug
	findByID.Description = input.Description
	findByID.Perks = input.Perks
	findByID.Price = input.Price
	findByID.UpdatedAt = time.Now()

	course, err := s.CourseRepository.Update(findByID)
	helper.PanicIfError(err)

	return formatter.ToCourseResponse(course), nil
}

func (s *CourseServiceImpl) FindByID(courseID int) (web.CourseResponse, error) {
	findByID, err := s.CourseRepository.FindByID(courseID)
	helper.PanicIfError(err)

	return formatter.ToCourseResponse(findByID), nil
}

func (s *CourseServiceImpl) FindByUserID(userID int) ([]web.CourseResponse, error) {
	courses, err := s.CourseRepository.FindByUserID(userID)
	helper.PanicIfError(err)

	return formatter.ToCourseResponses(courses), nil
}

func (s *CourseServiceImpl) FindAll() ([]web.CourseResponse, error) {
	courses, err := s.CourseRepository.FindAll()
	helper.PanicIfError(err)

	return formatter.ToCourseResponses(courses), nil
}

func (s *CourseServiceImpl) FindByCategory(category string) ([]web.CourseResponse, error) {
	courses, err := s.CourseRepository.FindByCategory(category)
	helper.PanicIfError(err)

	return formatter.ToCourseResponses(courses), nil
}

func (s *CourseServiceImpl) CountUserLearned(courseID int) (int, error) {
	countUsersLearned, err := s.CourseRepository.CountUsersLearned(courseID)
	helper.PanicIfError(err)

	return countUsersLearned, nil
}
