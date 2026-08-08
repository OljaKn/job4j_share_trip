package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/go-share-trip/internal/domain"
)

func (r *RepoPg) GetForUpdateByID(
	ctx context.Context,
	tx pgx.Tx,
	id uuid.UUID,
) (*domain.Trip, error) {
	var trip domain.Trip
	err := tx.QueryRow(ctx, `
		SELECT
			id,
			driver_id,
			from_point,
			to_point,
			departure_time,
			seats,
			status,
			created_at
		FROM trips
		WHERE id = $1 FOR UPDATE
	`, id).Scan(
		&trip.Id,
		&trip.DriverId,
		&trip.FromPoint,
		&trip.ToPoint,
		&trip.DepartureTime,
		&trip.Seats,
		&trip.Status,
		&trip.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get trip by id for update: %w", err)
	}
	return &trip, nil
}

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
