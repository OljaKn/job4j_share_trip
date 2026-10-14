package api

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/middleware"
	"job4j.ru/go-share-trip/internal/observability/logctx"
	"job4j.ru/go-share-trip/internal/service"
)

type CreateTripRequest struct {
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

	claims, err := middleware.ClaimsFromContext(c)
	if err != nil {
		return err
	}
	clientID := claims.Subject
	driverId, err := uuid.Parse(clientID)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid client_id token")
	}
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
	logger = logger.With(
		slog.String("client_id", clientID),
	)

	ctx = logctx.WithLogger(ctx, logger)

	cmd := service.CreateTripCommand{
		DriverId:      driverId,
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
