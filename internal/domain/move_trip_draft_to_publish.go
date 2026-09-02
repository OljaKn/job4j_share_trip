package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func MoveTripDraftToPublish(
	ctx context.Context,
	tx pgx.Tx,
	repo TripRepository,
	tripID, driverID uuid.UUID,
) (*Trip, bool, error) {
	trip, err := repo.GetForUpdateByID(ctx, tx, tripID)
	if err != nil {
		return nil, false, fmt.Errorf("tripRepository.GetForUpdateByID: %w", err)
	}

	if trip.DriverId != driverID {
		return nil, false, fmt.Errorf("%w: client %s is not driver of trip %s", ErrForbidden, driverID, tripID)
	}

	if trip.Status == Public {
		return trip, false, nil
	}

	if trip.Status != Draft {
		return nil, false, fmt.Errorf("%w: expected %s, got %s", ErrInvalidStatus, Draft, trip.Status)
	}

	fromStatus := trip.Status
	trip.Status = Public

	updatedTrip, err := repo.Update(ctx, tx, trip, fromStatus)
	if err != nil {
		return nil, false, fmt.Errorf("tripRepository.Update: %w", err)
	}

	return updatedTrip, true, nil
}
