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

	return c.Status(http.StatusOK).Send(nil)
}

func (s *Server) taskCreateHandler(c *fiber.Ctx) error {

	// проверка что в url нет лишних параметров
	if len(c.Queries()) > 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Title:  "Extra parameters in query url.",
			Status: http.StatusUnprocessableEntity,
			Detail: "Extra parameters in query url.",
		})
	}

	// получает тело запроса json и декодируем его в структуру
	var req models.TaskRequest
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResult{
			Title:  "Bad json.",
			Status: http.StatusBadRequest,
			Detail: "Bad json.",
		})
	}

	// валидация полей структуры
	validErrs := s.validator.Struct(req)
	var validationErrors validator.ValidationErrors
	errors.As(validErrs, &validationErrors)

	if len(validationErrors) > 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(models.ErrorResult{
			Title:  "Some error in json.",
			Status: http.StatusUnprocessableEntity,
			Detail: validationErrors.Error(),
		})
	}

	// вызываем метод логики
	res, err := s.service.AddTask(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResult{
			Title:  "Logic error.",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	}
	response := models.TaskAddResponse{}
	for _, val := range res {
		response.Tasks = append(response.Tasks, models.ID{Id: strconv.Itoa(val)})
	}

	resJson, _ := json.Marshal(response)

	return c.Status(http.StatusCreated).JSON(string(resJson))
}

func (s *Server) taskGetHandler(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).Send(nil)
}

func (s *Server) taskUpdateHandler(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).Send(nil)
}

func (s *Server) taskDeleteHandler(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).Send(nil)
}
