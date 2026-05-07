package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"job4j.ru/go-share-trip/internal/domain"
)

type RepoPg struct {
	pool *pgxpool.Pool
}

func NewRepoPg(pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{pool: pool}
}

type TripRepository interface {
	Ping(ctx context.Context) error
	Create(ctx context.Context, tr *domain.Trip) error
	Get(ctx context.Context, id uuid.UUID) (*domain.Trip, error)
}
