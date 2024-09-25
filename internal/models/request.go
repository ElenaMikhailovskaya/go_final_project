package models

type TaskRequest struct {
	Date    string `json:"date"`
	Title   string `json:"title" validate:"required"`
	Comment string `json:"comment" validate:"required"`
	Repeat  string `json:"repeat"`
}
