package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TripRepository interface {
	Create(ctx context.Context, tx pgx.Tx, trip *Trip) error
	Get(ctx context.Context, id uuid.UUID) (*Trip, error)
	Ping(ctx context.Context) error
	Update(ctx context.Context, tx pgx.Tx, trip *Trip, event OutboxEvent, fromStatus StatusTrip) (*Trip, error)
}
