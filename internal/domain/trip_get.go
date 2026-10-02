package domain

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"job4j.ru/go-share-trip/internal/observability/logctx"
)

func GetTrip(ctx context.Context, repo TripRepository, id uuid.UUID) (*Trip, error) {
	logger := logctx.Logger(ctx).With(
		slog.String("layer", "usecase"),
		slog.String("usecase", "TripUsecase.GetTrip"),
		slog.String("trip_id", id.String()),
	)
	logger.Info("get trip usecase started")

	trip, err := repo.Get(ctx, id)
	if err != nil {
		logger.Error(
			"repository get trip failed",
			slog.Any("error", err),
		)
		return nil, fmt.Errorf("repoTrip.Get: %w", err)
	}

	logger.Info("get trip usecase completed")
	return trip, nil
}
