package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"job4j.ru/go-share-trip/configs"
	"job4j.ru/go-share-trip/internal/api"
	"job4j.ru/go-share-trip/internal/app"
	"job4j.ru/go-share-trip/internal/middleware"
	"job4j.ru/go-share-trip/internal/observability/metrics"
	"job4j.ru/go-share-trip/internal/observability/tracing"
	"job4j.ru/go-share-trip/internal/repositories"
	"job4j.ru/go-share-trip/internal/service"
)

func main() {
	logger, logFile, err := app.NewLogger()
	if err != nil {
		log.Fatal("failed to init logger:", err)
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Printf("failed to close log file: %v", err)
		}
	}()
	configs.InitConfig()
	dbUrl := configs.GetDBConfig().DSN()

	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatal("connection fail:", err)
	}
	defer dbPool.Close()
	tp, err := tracing.NewProvider(ctx, tracing.Config{
		ServiceName:    "share-trip",
		ServiceVersion: "1.0.0",
		Environment:    "local",
		Endpoint:       "localhost:4319",
	})
	if err != nil {
		logger.Error("init tracing failed", "error", err)
		os.Exit(1)
	}

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer cancel()

		if err := tp.Shutdown(shutdownCtx); err != nil {
			logger.Error("shutdown tracing failed", "error", err)
		}
	}()
	registry := prometheus.NewRegistry()
	metric := metrics.New(registry)
	repo := repositories.NewRepoPg(metric, dbPool)
	outboxRepo := repositories.NewOutboxRepo(dbPool)
	service := service.NewTripService(metric, repo, outboxRepo, dbPool)
	handler := api.NewServer(service, registry, metric)
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(middleware.Correlation(logger))
	app.Use(middleware.NewHTTPMetricsMiddleware(metric))
	app.Use(tracing.NewFiberMiddleware())
	handler.Route(app.Group(""))
	port := configs.GetServerConfig()
	log.Fatal(app.Listen(":" + port.Port))

}
