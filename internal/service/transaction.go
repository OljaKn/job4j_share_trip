package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"job4j.ru/go-share-trip/internal/observability/logctx"
)

func tx[T any](
	ctx context.Context,
	pool *pgxpool.Pool,
	block func(tx pgx.Tx) (*T, error),
) (*T, error) {
	logger := logctx.Logger(ctx).With(
		slog.String("layer", "transaction"),
	)

	logger.Info("begin transaction")

	txBegin, err := pool.Begin(ctx)
	if err != nil {
		logger.Error(
			"failed to begin transaction",
			slog.Any("error", err),
		)
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		err := txBegin.Rollback(ctx)
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			logger.Error(
				"rollback transaction failed",
				slog.Any("error", err),
			)
		}
	}()

	res, err := block(txBegin)
	if err != nil {
		logger.Error(
			"transaction block failed",
			slog.Any("error", err),
		)
		return nil, fmt.Errorf("transaction block: %w", err)
	}

	if err = txBegin.Commit(ctx); err != nil {
		logger.Error(
			"failed to commit transaction",
			slog.Any("error", err),
		)
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	logger.Info("commit transaction")

	return res, nil
}
