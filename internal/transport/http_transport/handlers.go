package http_transport

import (
	"encoding/json"
	"errors"
	"github.com/ElenaMikhailovskaya/go_final_project/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func (s *Server) pingHandler(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).Send(nil)
}

func (s *Server) tasksListHandler(c *fiber.Ctx) error {
	// проверка что в url нет лишних параметров
	if len(c.Queries()) > 1 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Extra parameters in query url.",
		})
	}

	var query string
	query = c.Query("query")

	if query != "" {
		res, err := s.service.Search(query)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(models.ErrorResult{
				Error: err.Error(),
			})
		}
		return c.Status(http.StatusOK).JSON(res)
	} else {
		// вызываем метод логики
		res, err := s.service.GetTaskList()
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(models.ErrorResult{
				Error: err.Error(),
			})
		}
		return c.Status(http.StatusOK).JSON(res)
	}
}

func (s *Server) taskGetHandler(c *fiber.Ctx) error {

	// проверка что в url нет лишних параметров
	if len(c.Queries()) > 1 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Extra parameters in query url.",
		})
	}

	var taskId string
	taskId = c.Query("id")

	if taskId == "" {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Param id is empty.",
		})
	}

	// вызываем метод логики
	res, err := s.service.GetById(taskId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResult{
			Error: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (s *Server) taskCreateHandler(c *fiber.Ctx) error {

	// проверка что в url нет лишних параметров
	if len(c.Queries()) > 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Extra parameters in query url.",
		})
	}

	// получает тело запроса json и декодируем его в структуру
	var req models.TaskRequest
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResult{
			Error: "Bad json.",
		})
	}

	// валидация полей структуры
	validErrs := s.validator.Struct(req)
	var validationErrors validator.ValidationErrors
	errors.As(validErrs, &validationErrors)

	if len(validationErrors) > 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Some error in json.",
		})
	}

	// вызываем метод логики
	res, err := s.service.AddTask(req)
	if err != nil || res == 0 {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResult{
			Error: err.Error(),
		})
	}

	response := models.ID{Id: strconv.Itoa(res)}

	return c.Status(http.StatusCreated).JSON(response)
}

func (s *Server) taskDoneHandler(c *fiber.Ctx) error {

	// проверка что в url нет лишних параметров
	if len(c.Queries()) > 1 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Extra parameters in query url.",
		})
	}

	var taskId string
	taskId = c.Query("id")

	if taskId == "" {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Param id is empty.",
		})
	}

	// вызываем метод логики
	err := s.service.DoneTask(taskId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResult{
			Error: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Task{})
}

func (s *Server) taskUpdateHandler(c *fiber.Ctx) error {

	// проверка что в url нет лишних параметров
	if len(c.Queries()) > 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Extra parameters in query url.",
		})
	}

	// получает тело запроса json и декодируем его в структуру
	var req models.TaskUpdateRequest
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResult{
			Error: "Bad json.",
		})
	}

	// валидация полей структуры
	validErrs := s.validator.Struct(req)
	var validationErrors validator.ValidationErrors
	errors.As(validErrs, &validationErrors)

	if len(validationErrors) > 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Some error in json.",
		})
	}

	// вызываем метод логики
	err = s.service.UpdateTask(req, false)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResult{
			Error: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Task{})

}

func (s *Server) taskDeleteHandler(c *fiber.Ctx) error {

	// проверка что в url нет лишних параметров
	if len(c.Queries()) > 1 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Extra parameters in query url.",
		})
	}

	var taskId string
	taskId = c.Query("id")

	if taskId == "" {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Param id is empty.",
		})
	}

	// вызываем метод логики
	err := s.service.DeleteTask(taskId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResult{
			Error: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Task{})
}

func (s *Server) nextDateHandler(c *fiber.Ctx) error {

	// проверка что в url нет лишних параметров
	if len(c.Queries()) > 3 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: "Extra parameters in query url.",
		})
	}

	now := c.Query("now")
	date := c.Query("date")
	repeat := c.Query("repeat")

	nextDate, err := s.service.NextDate(now, date, repeat, false)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Error: err.Error(),
		})
	}

	return c.Status(http.StatusOK).Send([]byte(nextDate))
}
