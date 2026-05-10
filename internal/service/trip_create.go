package service

import (
	"context"
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

func (s *Service) CreateTrip(ctx context.Context, com CreateTripCommand) (*domain.Trip, error) {
	return domain.CreateTrip(
		ctx,
		s.Repository,
		com.DriverId,
		com.FromPoint,
		com.ToPoint,
		com.DepartureTime,
		com.Seats,
	)
}
