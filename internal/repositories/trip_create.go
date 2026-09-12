package repositories

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"job4j.ru/go-share-trip/internal/domain"
	"job4j.ru/go-share-trip/internal/observability/logctx"
)

func (r *RepoPg) Create(ctx context.Context, tx pgx.Tx, tr *domain.Trip) error {
	logger := logctx.Logger(ctx).With(
		slog.String("layer", "repository"),
		slog.String("repository", "TripRepository"),
		slog.String("operation", "Create"),
		slog.String("trip_id", tr.Id.String()),
		slog.String("client_id", tr.DriverId.String()),
	)

	logger.Info("insert trip started")
	_, err := tx.Exec(
		ctx,
		`INSERT INTO trips(id, driver_id, from_point, to_point, departure_time, seats, status, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		tr.Id, tr.DriverId, tr.FromPoint, tr.ToPoint, tr.DepartureTime, tr.Seats, tr.Status, tr.CreatedAt,
	)
	if err != nil {
		logger.Error(
			"insert trip failed",
			slog.Any("error", err),
		)
		return fmt.Errorf("tx.Exec create trip: %w", err)
	}

	logger.Info("insert trip completed")
	_, err = tx.Exec(
		ctx,
		`INSERT INTO trip_history(id, trip_id, to_status, created_at) VALUES($1, $2, $3, $4)`,
		uuid.New(), tr.Id, tr.Status, tr.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("tx.Exec insert history: %w", err)
	}
	return nil
}
