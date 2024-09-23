package http_transport

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (s *Server) pingHandler(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).Send(nil)
}

func (s *Server) tasksListHandler(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).Send(nil)
}

func (s *Server) taskCreateHandler(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).Send(nil)
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
