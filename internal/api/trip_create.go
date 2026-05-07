package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/service"
)

type CreateTripRequest struct {
	DriverId      uuid.UUID `json:"driver_id"`
	FromPoint     string    `json:"from_point"`
	ToPoint       string    `json:"to_point"`
	DepartureTime time.Time `json:"departure_time"`
	Seats         int       `json:"seats"`
}

type CreateTripResponse struct {
}

func (h *Handler) CreateTrip(c *fiber.Ctx) error {
	var req CreateTripRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.FromPoint == "" {
		return fiber.NewError(fiber.StatusBadRequest, "point of departure is required")
	}
	if req.ToPoint == "" {
		return fiber.NewError(fiber.StatusBadRequest, "point of arrival is required")
	}
	cmd := service.CreateTripCommand{
		DriverId:      req.DriverId,
		FromPoint:     req.FromPoint,
		ToPoint:       req.ToPoint,
		DepartureTime: req.DepartureTime,
		Seats:         req.Seats,
	}
	err := h.server.CreateTrip(c.Context(), cmd)
	if err != nil {
		log.Errorw("s.Repository.Create", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "created"})
}
