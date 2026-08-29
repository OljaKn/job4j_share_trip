package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"job4j.ru/go-share-trip/internal/domain"
)

type OutboxRepo struct {
	pool *pgxpool.Pool
}

func NewOutboxRepo(pool *pgxpool.Pool) *OutboxRepo {
	return &OutboxRepo{pool: pool}
}

func (r *OutboxRepo) Create(ctx context.Context, tx pgx.Tx, event domain.OutboxEvent) error {
	_, err := tx.Exec(ctx,
		`INSERT INTO outbox_event(id, event_name, aggregate_id, payload, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		event.Id, event.EventName, event.AggregateId, event.Payload, event.CreatedAt)
	if err != nil {
		return fmt.Errorf("insert outbox_event: %w", err)
	}
	return nil
}
