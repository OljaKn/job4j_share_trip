package repositories

import (
	"context"
	"errors"
	"fmt"

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
