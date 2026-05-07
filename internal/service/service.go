package service

import (
	"job4j.ru/go-share-trip/internal/repositories"
)

type Service struct {
	Repository *repositories.RepoPg
}

func NewServer(repo *repositories.RepoPg) *Service {
	return &Service{Repository: repo}
}
