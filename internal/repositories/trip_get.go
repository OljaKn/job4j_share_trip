package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/domain"
)

func (r *RepoPg) Get(ctx context.Context, id uuid.UUID) (*domain.Trip, error) {
	var trip domain.Trip
	err := r.pool.QueryRow(
		ctx,
		`select id, driver_id, from_point, to_point, departure_time, seats, status, created_at from trips where id = $1`,
		id,
	).Scan(&trip.Id, &trip.DriverId, &trip.FromPoint, &trip.ToPoint, &trip.DepartureTime, &trip.Seats, &trip.Status, &trip.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return &trip, nil
}
