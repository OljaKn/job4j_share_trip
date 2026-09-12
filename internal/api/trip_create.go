package api

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/observability/logctx"
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
	ID            uuid.UUID `json:"id"`
	DriverId      uuid.UUID `json:"driver_id"`
	FromPoint     string    `json:"from_point"`
	ToPoint       string    `json:"to_point"`
	DepartureTime time.Time `json:"departure_time"`
	Seats         int       `json:"seats"`
	Status        string    `json:"trip_status"`
	CreatedAt     time.Time `json:"created_at"`
}

func (h *Server) CreateTrip(c *fiber.Ctx) error {
	ctx := c.UserContext()

	logger := logctx.Logger(ctx).With(
		slog.String("server", "TripServer"),
		slog.String("handler", "CreateTrip"),
	)
	var req CreateTripRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Warn(
			"create trip failed: invalid json body",
			slog.Any("error", err),
		)
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.DriverId == uuid.Nil {
		logger.Warn("create trip failed: client_id is required")
		return fiber.NewError(fiber.StatusBadRequest, "driver_id is required")
	}
	if req.FromPoint == "" {
		logger.Warn("create trip failed: fromPoint is required")
		return fiber.NewError(fiber.StatusBadRequest, "point of departure is required")
	}
	if req.ToPoint == "" {
		logger.Warn("create trip failed: toPoint is required")
		return fiber.NewError(fiber.StatusBadRequest, "point of arrival is required")
	}
	logger = logger.With(
		slog.String("client_id", req.DriverId.String()),
	)

	ctx = logctx.WithLogger(ctx, logger)

	logger.Info("create trip request accepted")
	cmd := service.CreateTripCommand{
		DriverId:      req.DriverId,
		FromPoint:     req.FromPoint,
		ToPoint:       req.ToPoint,
		DepartureTime: req.DepartureTime,
		Seats:         req.Seats,
	}
	trip, err := h.server.CreateTrip(ctx, cmd)
	if err != nil {
		logger.Error(
			"create trip failed",
			slog.Any("error", err),
		)
		//log.Errorw("s.Repository.Create", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	logger.Info(
		"create trip completed",
		slog.String("trip_id", trip.Id.String()),
	)
	return c.Status(fiber.StatusCreated).JSON(CreateTripResponse{
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
