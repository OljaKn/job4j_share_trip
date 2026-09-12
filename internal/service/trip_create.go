package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/go-share-trip/internal/domain"
	"job4j.ru/go-share-trip/internal/observability/logctx"
)

type CreateTripCommand struct {
	DriverId      uuid.UUID `json:"driver_id"`
	FromPoint     string    `json:"fromPoint"`
	ToPoint       string    `json:"toPoint"`
	DepartureTime time.Time `json:"departureTime"`
	Seats         int       `json:"seats"`
}

func (s *Service) CreateTrip(ctx context.Context, com CreateTripCommand) (*domain.Trip, error) {
	logger := logctx.Logger(ctx).With(
		slog.String("service", "TripService"),
		slog.String("operation", "CreateTrip"),
		slog.String("client_id", com.DriverId.String()),
	)

	logger.Info("create trip started")
	res, err := tx(ctx, s.pool, func(tx pgx.Tx) (*domain.Trip, error) {
		txLogger := logger.With(
			slog.String("layer", "transaction"),
		)

		txLogger.Info("transaction started")
		trip, err := domain.CreateTrip(
			ctx,
			tx,
			s.Repository,
			com.DriverId,
			com.FromPoint,
			com.ToPoint,
			com.DepartureTime,
			com.Seats,
		)
		if err != nil {
			txLogger.Error(
				"create trip usecase failed",
				slog.Any("error", err),
			)
			return nil, fmt.Errorf("usecase.CreateTrip: %w", err)
		}
		txLogger.Info(
			"transaction completed",
			slog.String("trip_id", trip.Id.String()),
		)
		return trip, nil
	})
	if err != nil {
		logger.Error(
			"create trip failed",
			slog.Any("error", err),
		)
		return nil, fmt.Errorf("failed in transaction: %w", err)
	}
	logger.Info(
		"create trip completed",
		slog.String("trip_id", res.Id.String()),
	)
	return res, err
}
