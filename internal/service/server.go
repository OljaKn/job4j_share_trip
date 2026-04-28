package service

import (
	"context"

	"job4j.ru/go-share-trip/internal/storage"
)

type Server struct {
	Repository *storage.RepoPg
}

func NewServer(repo *storage.RepoPg) *Server {
	return &Server{Repository: repo}
}

func (s *Server) CheckReady(ctx context.Context) error {
	return s.Repository.Ping(ctx)
}
