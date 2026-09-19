package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"job4j.ru/go-share-trip/internal/observability/metrics"
)

type RepoPg struct {
	metrics *metrics.Metrics
	pool    *pgxpool.Pool
}

func NewRepoPg(metrics *metrics.Metrics, pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{
		pool:    pool,
		metrics: metrics,
	}
}
