package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func (h *Server) GetTrip(c *fiber.Ctx) error {
	tracer := otel.Tracer("trip-api")

	ctx, span := tracer.Start(c.UserContext(), "GetTripHandler")
	defer span.End()

	c.Set("trace-id", span.SpanContext().TraceID().String())
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
