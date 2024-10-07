package http_transport

import (
	"errors"
	"github.com/ElenaMikhailovskaya/go_final_project/internal/models"
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New(service taskInterfaceMethods) (*Server, error) {
	var cfg Cfg
	e := env.Parse(&cfg)
	if e != nil {
		return nil, e
	}

	s := new(Server)
	s.cfg = cfg

	s.service = service
	if service == nil {
		return nil, errors.New("param `service` is required")
	}

	s.validator = validator.New(validator.WithRequiredStructEnabled())
	s.a = NewFiberApp()

	return s.SetupHandlers(), nil
}

func (s *Server) SetupHandlers() *Server {
	v1 := s.a.Group(models.ApiPath)

	s.a.Static(webDir, "web")                          // фронт
	v1.Get(models.ApiPathPing, s.pingHandler)          // пинг приложения
	v1.Get(models.ApiPathTasks, s.tasksListHandler)    // Получение списка задач
	v1.Post(models.ApiPathTask, s.taskCreateHandler)   // Добавление задачи
	v1.Post(models.ApiTaskDone, s.taskDoneHandler)     // Выполнение задачи
	v1.Get(models.ApiPathTask, s.taskGetHandler)       // Получение задачи
	v1.Put(models.ApiPathTask, s.taskUpdateHandler)    // Изменение задачи
	v1.Delete(models.ApiPathTask, s.taskDeleteHandler) // Удаление задачи
	v1.Get(models.ApiNextDate, s.nextDateHandler)      // След дата

	return s
}

func NewFiberApp() *fiber.App {
	a := fiber.New(fiber.Config{
		ServerHeader: "Fiber", // добавляем заголовок для идентификации сервера
	})

	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	a.Use(recover.New())

	return a
}
