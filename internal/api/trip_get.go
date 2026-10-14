package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"job4j.ru/go-share-trip/internal/middleware"
)

func (h *Server) GetTrip(c *fiber.Ctx) error {
	tracer := otel.Tracer("trip-api")

	ctx, span := tracer.Start(c.UserContext(), "GetTripHandler")
	defer span.End()

	c.Set("trace-id", span.SpanContext().TraceID().String())
	claims, err := middleware.ClaimsFromContext(c)
	if err != nil {
		return err
	}

	userID := claims.Subject
	_, err = uuid.Parse(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid user ID in token")
	}
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}
	tripId, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid format id")
	}
	trip, err := h.server.GetTrip(ctx, tripId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	span.SetAttributes(
		attribute.String("trip_id", tripId.String()),
		attribute.String("driver_id", trip.DriverId.String()),
	)
	return c.Status(fiber.StatusOK).JSON(trip)
}
