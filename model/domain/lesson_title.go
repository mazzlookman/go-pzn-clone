package domain

import "time"

type LessonTitle struct {
	ID        int
	CourseID  int
	InOrder   int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
