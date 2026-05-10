package service

import (
	"job4j.ru/go-share-trip/internal/domain"
)

type Service struct {
	Repository domain.TripRepository
}

func NewServer(repo domain.TripRepository) *Service {
	return &Service{Repository: repo}
}
