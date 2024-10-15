package models

type TaskRequest struct {
	Date    string `json:"date"`
	Title   string `json:"title" validate:"required"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}
type TaskUpdateRequest struct {
	Id      string `json:"id" validate:"required,numeric"`
	Date    string `json:"date" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}
type Search struct {
	Query string
	Date  string
}
