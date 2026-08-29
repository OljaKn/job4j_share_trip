package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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
		trip, statusChanged, err := domain.MoveTripDraftToPublish(
			ctx,
			tx,
			s.Repository,
			com.TripId,
			com.DriverId,
		)
		if err != nil {
			return nil, err
		}

		if !statusChanged {
			return trip, nil
		}

		payload, err := json.Marshal(domain.TripPublishedPayload{
			TripID: trip.Id.String(),
		})
		if err != nil {
			return nil, fmt.Errorf("marshal payload: %w", err)
		}

		event := domain.OutboxEvent{
			Id:          uuid.New(),
			EventName:   domain.TripPublished,
			AggregateId: trip.Id,
			Payload:     string(payload),
			CreatedAt:   time.Now(),
		}

		err = s.OutboxRepository.Create(ctx, tx, event)
		if err != nil {
			return nil, fmt.Errorf("outboxRepository.Create: %w", err)
		}
		return trip, nil
	})
}
