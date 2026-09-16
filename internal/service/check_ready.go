package service

import (
	"context"
)

func (s *Service) CheckReady(ctx context.Context) error {
	return s.Repository.Ping(ctx)
}
