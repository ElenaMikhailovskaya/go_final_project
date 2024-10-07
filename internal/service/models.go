package service

import "github.com/ElenaMikhailovskaya/go_final_project/internal/models"

type Server struct {
	DBase DB
}

type DB interface {
	CreateTable()
	Close()
	Add(task models.TaskRequest, nextDate string) (int, error)
	GetList() (models.TasksGetResponse, error)
	GetById(id string) (models.Task, error)
	Update(task models.TaskUpdateRequest) error
	Delete(id string) error
	Search(search models.Search) (models.TasksGetResponse, error)
}

const (
	MaxMoveDays = 400
)
