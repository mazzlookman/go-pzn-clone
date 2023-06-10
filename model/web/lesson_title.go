package web

type LessonTitleInput struct {
	CourseID int    `json:"course_id"`
	InOrder  int    `json:"in_order"`
	Title    string `json:"title"`
}

type LessonTitleResponse struct {
	ID       int    `json:"id"`
	CourseID int    `json:"course_id"`
	InOrder  int    `json:"in_order"`
	Title    string `json:"title"`
}
