package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) GetTrip(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}
	tripId, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid format id")
	}
	trip, err := h.server.GetTrip(c.Context(), tripId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	return c.Status(fiber.StatusOK).JSON(trip)
}
