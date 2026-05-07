package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/domain"
)

func (r *RepoPg) Create(ctx context.Context, tr *domain.Trip) error {
	_, err := r.pool.Exec(
		ctx,
		`insert into trips(id, driver_id, from_point, to_point, departure_time, seats, status, created_at) values($1, $2, $3, $4, $5, $6, $7, $8)`,
		tr.Id, tr.DriverId, tr.FromPoint, tr.ToPoint, tr.DepartureTime, tr.Seats, tr.Status, tr.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}
	_, err1 := r.pool.Exec(
		ctx,
		`insert into trip_history(id, trip_id, from_status, to_status, created_at) values($1, $2, $3, $4, $5)`,
		uuid.New(), tr.Id, nil, tr.Status, tr.CreatedAt,
	)
	if err1 != nil {
		return fmt.Errorf("r.pool.Exec: %w", err1)
	}
	return nil
}
