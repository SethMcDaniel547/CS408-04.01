package models

type Course struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"course_code"`
}

type Assignment struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	DueAt          string  `json:"due_at"`
	PointsPossible float64 `json:"points_possible"`
}
