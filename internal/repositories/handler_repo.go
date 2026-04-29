package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepoPg struct {
	pool *pgxpool.Pool
}

func NewRepoPg(pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{pool: pool}
}

func (r *RepoPg) Ping(ctx context.Context) error {
	var res string
	err := r.pool.QueryRow(
		ctx,
		`select version()`).Scan(&res)
	return err
}
