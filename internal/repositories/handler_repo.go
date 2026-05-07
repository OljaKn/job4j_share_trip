package repositories

import (
	"context"
)

func (r *RepoPg) Ping(ctx context.Context) error {
	var res string
	err := r.pool.QueryRow(
		ctx,
		`select version()`).Scan(&res)
	return err
}
