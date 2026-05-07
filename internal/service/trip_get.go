package service

import (
	"context"

	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/domain"
)

func (s *Service) GetTrip(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	trip, err := s.Repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return trip, nil
}
