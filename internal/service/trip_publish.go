package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/go-share-trip/internal/domain"
)

type PublishTripCommand struct {
	TripId   uuid.UUID
	DriverId uuid.UUID
}

func (s *Service) PublishTrip(ctx context.Context, com PublishTripCommand) (*domain.Trip, error) {
	return tx(ctx, s.pool, func(tx pgx.Tx) (*domain.Trip, error) {
		return domain.MoveTripDraftToPublish(
			ctx,
			tx,
			s.Repository,
			com.TripId,
			com.DriverId,
		)
	})
}
