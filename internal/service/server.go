package service

import (
	"context"

	"job4j.ru/go-share-trip/internal/repositories"
)

type Server struct {
	Repository *repositories.RepoPg
}

func NewServer(repo *repositories.RepoPg) *Server {
	return &Server{Repository: repo}
}

func (s *Server) CheckReady(ctx context.Context) error {
	return s.Repository.Ping(ctx)
}
