package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func MoveTripDraftToPublish(
	ctx context.Context,
	tx pgx.Tx,
	repo TripRepository,
	tripID, driverID uuid.UUID,
) (*Trip, error) {
	trip, err := repo.Get(ctx, tripID)
	if err != nil {
		return nil, fmt.Errorf("tripRepository.Get: %w", err)
	}

	if trip.DriverId != driverID {
		return nil, fmt.Errorf("%w: client %s is not driver of trip %s", ErrForbidden, driverID, tripID)
	}

	if trip.Status == Public {
		return trip, nil
	}

	if trip.Status != Draft {
		return nil, fmt.Errorf("%w: expected %s, got %s", ErrInvalidStatus, Draft, trip.Status)
	}

	fromStatus := trip.Status
	trip.Status = Public
	payload := fmt.Sprintf(`{"trip_id": "%s"}`, tripID)
	event := OutboxEvent{
		Id:          uuid.New(),
		EventName:   "event_published",
		AggregateId: tripID,
		Payload:     payload,
		CreatedAt:   time.Now(),
	}

	updatedTrip, err := repo.Update(ctx, tx, trip, event, fromStatus)
	if err != nil {
		return nil, fmt.Errorf("tripRepository.Update: %w", err)
	}

	return updatedTrip, nil
}
