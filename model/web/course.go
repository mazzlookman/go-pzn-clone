package web

type CourseInput struct {
	AuthorID    int    `json:"author_id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Perks       string `json:"perks"`
	Price       int    `json:"price"`
	Banner      string `json:"banner"`
}

type CourseResponse struct {
	ID          int    `json:"id"`
	AuthorID    int    `json:"author_id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Perks       string `json:"perks"`
	Price       int    `json:"price"`
	Banner      string `json:"banner"`
}
