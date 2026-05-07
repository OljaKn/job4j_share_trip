package domain

import (
	"time"

	"github.com/google/uuid"
)

func NewTrip(driverId uuid.UUID, fromPoint string, toPoint string, departureTime time.Time, seats int) (*Trip, error) {
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
