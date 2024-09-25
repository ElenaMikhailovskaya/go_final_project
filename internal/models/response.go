package models

type TaskAddResponse struct {
	Tasks []ID `json:""`
}

type ID struct {
	Id string `json:"id"`
}
