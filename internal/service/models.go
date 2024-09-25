package service

import "github.com/ElenaMikhailovskaya/go_final_project/internal/models"

type Server struct {
	DBase DB
}

type DB interface {
	CreateTable()
	Close()
	Add(task models.TaskRequest, nextDate string) (int, error)
}

const (
	MaxMoveDays = 400
)
