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
	fromStatus domain.StatusTrip,
) (*domain.Trip, error) {
	_, err := tx.Exec(ctx,
		`UPDATE trips SET status = $1 WHERE id = $2`,
		trip.Status, trip.Id)
	if err != nil {
		return nil, fmt.Errorf("update trip: %w", err)
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO trip_history(id, trip_id, from_status, to_status, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		uuid.New(), trip.Id, fromStatus, trip.Status, time.Now())
	if err != nil {
		return nil, fmt.Errorf("insert history: %w", err)
	}

	return trip, nil
}
