package api

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Route(route fiber.Router) {
	route.Get("/ready", h.Ready)
}
