package formatter

import (
	"go-pzn-clone/model/domain"
	"go-pzn-clone/model/web"
)

func ToUserResponse(user domain.User, jwtToken string) web.UserResponse {
	return web.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
		Token:  jwtToken,
	}
}

func ToCourseResponse(course domain.Course) web.CourseResponse {
	return web.CourseResponse{
		ID:          course.ID,
		AuthorID:    course.AuthorID,
		Title:       course.Title,
		Slug:        course.Slug,
		Description: course.Description,
		Perks:       course.Perks,
		Price:       course.Price,
		Banner:      course.Banner,
	}
}

func ToCourseResponses(courses []domain.Course) []web.CourseResponse {
	courseResponses := []web.CourseResponse{}

	for _, course := range courses {
		courseResponse := ToCourseResponse(course)
		courseResponses = append(courseResponses, courseResponse)
	}

	return courseResponses
}

func ToLessonTitlesResponse(title domain.LessonTitle) web.LessonTitleResponse {
	return web.LessonTitleResponse{
		ID:       title.ID,
		CourseID: title.CourseID,
		InOrder:  title.InOrder,
		Title:    title.Title,
	}
}

func ToLessonTitlesResponses(titles []domain.LessonTitle) []web.LessonTitleResponse {
	lessonTitleResponses := []web.LessonTitleResponse{}
	for _, lessonTitle := range titles {
		lessonTitleResponse := ToLessonTitlesResponse(lessonTitle)
		lessonTitleResponses = append(lessonTitleResponses, lessonTitleResponse)
	}

	return lessonTitleResponses
}

func ToLessonContentResponse(content domain.LessonContent) web.LessonContentResponse {
	return web.LessonContentResponse{
		ID:            content.ID,
		LessonTitleID: content.LessonTitleID,
		InOrder:       content.InOrder,
		Content:       content.Content,
		Duration:      content.Duration,
	}
}

func ToLessonContentResponses(contents []domain.LessonContent) []web.LessonContentResponse {
	lessonContentResponses := []web.LessonContentResponse{}

	for _, lessonContent := range contents {
		lessonContentResponse := ToLessonContentResponse(lessonContent)
		lessonContentResponses = append(lessonContentResponses, lessonContentResponse)
	}

	return lessonContentResponses
}
