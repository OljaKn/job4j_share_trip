package service

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"job4j.ru/go-share-trip/internal/domain"
)

func (s *Service) GetTrip(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	ctx, span := otel.Tracer("TripService").Start(ctx, "TripService.GetTrip")
	defer span.End()
	trip, err := domain.GetTrip(ctx, s.Repository, id)
	if err != nil {
		return nil, err
	}
	return trip, nil
}
