package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	return tx(ctx, s.pool, func(tx pgx.Tx) (*domain.Trip, error) {
		return domain.CreateTrip(
			ctx,
			tx,
			s.Repository,
			com.DriverId,
			com.FromPoint,
			com.ToPoint,
			com.DepartureTime,
			com.Seats,
		)
	})
}
