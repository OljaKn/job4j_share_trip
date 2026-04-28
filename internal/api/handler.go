package api

import (
	"github.com/gofiber/fiber/v2"
	"job4j.ru/go-share-trip/internal/service"
)

type Handler struct {
	server *service.Server
}

func NewHandler(server *service.Server) *Handler {
	return &Handler{server: server}
}

func (h *Handler) Ready(c *fiber.Ctx) error {
	err := h.server.CheckReady(c.Context())
	if err != nil {
		return c.Status(500).SendString("Fail")
	}
	return c.SendString("Ok")
}
