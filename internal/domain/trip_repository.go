package domain

import (
	"context"

	"github.com/google/uuid"
)

type TripRepository interface {
	Create(ctx context.Context, trip *Trip) error
	Get(ctx context.Context, id uuid.UUID) (*Trip, error)
	Ping(ctx context.Context) error
}
