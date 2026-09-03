package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/domain"
	"job4j.ru/go-share-trip/internal/service"
)

type PublishTripRequest struct {
	TripId   uuid.UUID `json:"trip_id"`
	DriverId uuid.UUID `json:"driver_id"`
}

func (h *Server) PublishTrip(c *fiber.Ctx) error {
	var req PublishTripRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.TripId == uuid.Nil {
		return fiber.NewError(fiber.StatusBadRequest, "trip_id is required")
	}
	if req.DriverId == uuid.Nil {
		return fiber.NewError(fiber.StatusBadRequest, "driver_id is required")
	}

	trip, err := h.server.PublishTrip(c.Context(), service.PublishTripCommand{
		TripId:   req.TripId,
		DriverId: req.DriverId,
	})
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "trip not found")
		}
		if errors.Is(err, domain.ErrForbidden) {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}
		if errors.Is(err, domain.ErrInvalidStatus) {
			return fiber.NewError(fiber.StatusConflict, "invalid trip status")
		}
		log.Errorw("PublishTrip", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(trip)
}
