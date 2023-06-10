package domain

import "time"

type LessonContent struct {
	ID            int
	LessonTitleID int
	InOrder       int
	Content       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
