package service

import (
	"job4j.ru/go-share-trip/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Repository       domain.TripRepository
	OutboxRepository domain.OutboxRepository
	pool             *pgxpool.Pool
}

func NewServer(repo domain.TripRepository, event domain.OutboxRepository, pool *pgxpool.Pool) *Service {
	return &Service{Repository: repo, OutboxRepository: event, pool: pool}
}
