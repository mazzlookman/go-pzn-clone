package web

type LessonContentInput struct {
	LessonTitleID int `form:"lt_id"`
	InOrder       int `form:"in_order"`
	Content       string
	Duration      string
}

type LessonContentResponse struct {
	ID            int    `json:"id"`
	LessonTitleID int    `json:"lesson_title_id"`
	InOrder       int    `json:"in_order"`
	Content       string `json:"content"`
	Duration      string `json:"duration"`
}

type LessonContentIDFromURI struct {
	ID int `uri:"lc_id"`
}
