package api

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Ready(c *fiber.Ctx) error {
	err := h.server.CheckReady(c.Context())
	if err != nil {
		return c.Status(500).SendString("Fail")
	}
	return c.SendString("Ok")
}
