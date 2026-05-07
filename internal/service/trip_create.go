package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/domain"
)

type CreateTripCommand struct {
	DriverId      uuid.UUID `json:"driver_id"`
	FromPoint     string    `json:"fromPoint"`
	ToPoint       string    `json:"toPoint"`
	DepartureTime time.Time `json:"departureTime"`
	Seats         int       `json:"seats"`
}

func (s *Service) CreateTrip(ctx context.Context, com CreateTripCommand) error {
	if com.Seats <= 0 {
		return errors.New("incorrect number of seats")
	}
	if com.DepartureTime.Before(time.Now()) {
		return errors.New("the trip time must be in the future")
	}
	trip, err1 := domain.NewTrip(
		com.DriverId,
		com.FromPoint,
		com.ToPoint,
		com.DepartureTime,
		com.Seats,
	)
	if err1 != nil {
		return err1
	}
	err := s.Repository.Create(ctx, trip)
	if err != nil {
		return err
	}
	return nil
}
