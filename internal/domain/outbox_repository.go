package domain

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type OutboxRepository interface {
	Create(ctx context.Context, tx pgx.Tx, event OutboxEvent) error
}
