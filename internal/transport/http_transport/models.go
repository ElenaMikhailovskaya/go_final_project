package http_transport

import (
	"github.com/ElenaMikhailovskaya/go_final_project/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const webDir = "/"

type Server struct {
	a         *fiber.App
	validator *validator.Validate
	cfg       Cfg

	service taskInterfaceMethods
}

type taskInterfaceMethods interface {
	AddTask(task models.TaskRequest) (int, error)
	UpdateTask(req models.TaskUpdateRequest) error
	DeleteTask(id string) error
	GetTaskList() (models.TasksGetResponse, error)
	DoneTask(id string) error
	NextDate(now string, date string, repeat string) (string, error)
	GetById(id string) (models.Task, error)
	Search(query string) (models.TasksGetResponse, error)
}

type Cfg struct {
	Host string `env:"TODO_PORT" envDefault:"localhost:7540"`
}
