package models

type Task struct {
	ID          int    `json:"id"`
	Deadline    string `json:"deadline"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
