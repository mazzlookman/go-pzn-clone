package domain

import "time"

type Course struct {
	ID          int
	AuthorID    int
	Title       string
	Slug        string
	Description string
	Perks       string
	Price       int
	Banner      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
