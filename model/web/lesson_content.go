package web

type LessonContentInput struct {
	LessonTitleID int    `json:"lesson_title_id"`
	InOrder       int    `json:"in_order"`
	Content       string `json:"content"`
}

type LessonContentResponse struct {
	ID            int    `json:"id"`
	LessonTitleID int    `json:"lesson_title_id"`
	InOrder       int    `json:"in_order"`
	Content       string `json:"content"`
}
