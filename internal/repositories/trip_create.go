package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/go-share-trip/internal/domain"
)

func (r *RepoPg) Create(ctx context.Context, tx pgx.Tx, tr *domain.Trip) error {
	_, err := tx.Exec(
		ctx,
		`INSERT INTO trips(id, driver_id, from_point, to_point, departure_time, seats, status, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		tr.Id, tr.DriverId, tr.FromPoint, tr.ToPoint, tr.DepartureTime, tr.Seats, tr.Status, tr.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("tx.Exec insert trip: %w", err)
	}
	_, err = tx.Exec(
		ctx,
		`INSERT INTO trip_history(id, trip_id, from_status, to_status, created_at) VALUES($1, $2, $3, $4, $5)`,
		uuid.New(), tr.Id, nil, tr.Status, tr.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("tx.Exec insert history: %w", err)
	}
	return nil
}
