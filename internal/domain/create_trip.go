package domain

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/go-share-trip/internal/observability/logctx"
)

func NewTrip(driverId uuid.UUID, fromPoint string, toPoint string, departureTime time.Time, seats int) (*Trip, error) {
	if seats <= 0 {
		return nil, errors.New("incorrect number of seats")
	}
	if departureTime.Before(time.Now()) {
		return nil, errors.New("the trip time must be in the future")
	}

	return &Trip{
		Id:            uuid.New(),
		DriverId:      driverId,
		FromPoint:     fromPoint,
		ToPoint:       toPoint,
		DepartureTime: departureTime,
		Seats:         seats,
		Status:        Draft,
		CreatedAt:     time.Now(),
	}, nil
}

func CreateTrip(ctx context.Context, tx pgx.Tx, repo TripRepository, driverId uuid.UUID, fromPoint string, toPoint string, departureTime time.Time, seats int) (*Trip, error) {
	logger := logctx.Logger(ctx).With(
		slog.String("layer", "usecase"),
		slog.String("usecase", "TripUsecase.CreateTrip"),
		slog.String("client_id", driverId.String()),
	)
	logger.Info("create trip usecase started")
	trip, err := NewTrip(driverId, fromPoint, toPoint, departureTime, seats)
	if err != nil {
		logger.Warn(
			"create trip validation failed",
			slog.Any("error", err),
		)
		return nil, err
	}

	err = repo.Create(ctx, tx, trip)
	if err != nil {
		logger.Error(
			"repository create trip failed",
			slog.Any("error", err),
		)
		return nil, fmt.Errorf("repoTrip.Create: %w", err)
	}

	logger.Info(
		"create trip usecase completed",
		slog.String("trip_id", trip.Id.String()),
	)
	return trip, nil
}
