package api

import (
	"job4j.ru/go-share-trip/internal/service"
)

type Handler struct {
	server *service.Service
}

func NewHandler(server *service.Service) *Handler {
	return &Handler{
		server: server,
	}
}
