package service

import (
	"job4j.ru/go-share-trip/internal/domain"
	"job4j.ru/go-share-trip/internal/observability/metrics"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	metrics          *metrics.Metrics
	Repository       domain.TripRepository
	OutboxRepository domain.OutboxRepository
	pool             *pgxpool.Pool
}

func NewTripService(metrics *metrics.Metrics, repo domain.TripRepository, event domain.OutboxRepository, pool *pgxpool.Pool) *Service {
	return &Service{
		metrics:          metrics,
		Repository:       repo,
		OutboxRepository: event,
		pool:             pool}
}
