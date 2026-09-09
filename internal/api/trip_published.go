package api

import (
	"errors"
	"time"

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
type PublishTripResponse struct {
	ID            uuid.UUID `json:"id"`
	DriverId      uuid.UUID `json:"driver_id"`
	FromPoint     string    `json:"from_point"`
	ToPoint       string    `json:"to_point"`
	DepartureTime time.Time `json:"departure_time"`
	Seats         int       `json:"seats"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
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
		if errors.Is(err, domain.ErrConflict) {
			return fiber.NewError(fiber.StatusConflict, "invalid trip status")
		}
		if errors.Is(err, domain.ErrTripAlreadyPublished) {
			return fiber.NewError(fiber.StatusNoContent, "trip already published")
		}
		log.Errorw("PublishTrip", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(PublishTripResponse{
		ID:            trip.Id,
		DriverId:      trip.DriverId,
		FromPoint:     trip.FromPoint,
		ToPoint:       trip.ToPoint,
		DepartureTime: trip.DepartureTime,
		Seats:         trip.Seats,
		Status:        string(trip.Status),
		CreatedAt:     trip.CreatedAt,
	})
}
