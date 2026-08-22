package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/go-share-trip/internal/domain"
)

func (r *RepoPg) Update(
	ctx context.Context,
	tx pgx.Tx,
	trip *domain.Trip,
	event domain.OutboxEvent,
	fromStatus domain.StatusTrip,
) (*domain.Trip, error) {
	res, err := tx.Exec(ctx,
		`UPDATE trips SET status = $1, version = version + 1 WHERE id = $2 and version = $3`,
		trip.Status, trip.Id, trip.Version)
	if err != nil {
		return nil, fmt.Errorf("update trip: %w", err)
	}
	if res.RowsAffected() == 0 {
		return nil, fmt.Errorf("failed to update the trip")
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO outbox_event(id, event_name, aggregate_id, payload, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		event.Id, event.EventName, event.AggregateId, event.Payload, event.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("insert outbox_event: %w", err)
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO trip_history(id, trip_id, from_status, to_status, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		uuid.New(), trip.Id, fromStatus, trip.Status, time.Now())
	if err != nil {
		return nil, fmt.Errorf("insert history: %w", err)
	}
	trip.Version = trip.Version + 1

	return trip, nil
}
