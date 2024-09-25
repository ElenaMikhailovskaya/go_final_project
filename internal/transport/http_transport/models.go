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
	AddTask(task models.TaskRequest) ([]int, error)
	UpdateTask()
	DeleteTask()
	GetTaskList()
	GetTask()
}

type Cfg struct {
	Host string `env:"TODO_PORT" envDefault:"localhost:7540"`
}
