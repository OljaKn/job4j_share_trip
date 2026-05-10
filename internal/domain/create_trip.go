package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
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

func CreateTrip(ctx context.Context, repo TripRepository, driverId uuid.UUID, fromPoint string, toPoint string, departureTime time.Time, seats int) (*Trip, error) {
	trip, err := NewTrip(driverId, fromPoint, toPoint, departureTime, seats)
	if err != nil {
		return nil, err
	}

	err = repo.Create(ctx, trip)
	if err != nil {
		return nil, err
	}
	return trip, nil
}
