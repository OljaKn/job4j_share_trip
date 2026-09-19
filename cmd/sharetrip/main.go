package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"job4j.ru/go-share-trip/configs"
	"job4j.ru/go-share-trip/internal/api"
	"job4j.ru/go-share-trip/internal/app"
	"job4j.ru/go-share-trip/internal/middleware"
	"job4j.ru/go-share-trip/internal/observability/metrics"
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
	registry := prometheus.NewRegistry()
	m := metrics.New(registry)
	repo := repositories.NewRepoPg(m, dbPool)
	outboxRepo := repositories.NewOutboxRepo(dbPool)
	service := service.NewTripService(m, repo, outboxRepo, dbPool)
	handler := api.NewServer(service, registry, m)
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(middleware.Correlation(logger))
	app.Use(api.NewHTTPMetricsMiddleware(m))
	handler.Route(app.Group(""))
	port := configs.GetServerConfig()
	log.Fatal(app.Listen(":" + port.Port))

}
