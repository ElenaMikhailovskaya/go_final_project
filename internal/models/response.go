package models

type TaskAddResponse struct {
	Tasks []ID `json:""`
}

type ID struct {
	Id string `json:"id"`
}

type TasksGetResponse struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Title   string `json:"title,omitempty"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
}
