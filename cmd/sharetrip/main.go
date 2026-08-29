package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"job4j.ru/go-share-trip/configs"
	"job4j.ru/go-share-trip/internal/api"
	"job4j.ru/go-share-trip/internal/repositories"
	"job4j.ru/go-share-trip/internal/service"
)

func main() {
	configs.InitConfig()
	dbUrl := configs.GetDBConfig().DSN()

	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatal("connection fail:", err)
	}
	defer dbPool.Close()

	repo := repositories.NewRepoPg(dbPool)
	outboxRepo := repositories.NewOutboxRepo(dbPool)
	service := service.NewServer(repo, outboxRepo, dbPool)
	handler := api.NewHandler(service)
	app := fiber.New()
	handler.Route(app)

	log.Fatal(app.Listen(":8080"))
}
