package api

import (
	"job4j.ru/go-share-trip/internal/service"
)

type Server struct {
	server *service.Service
}

func NewServer(server *service.Service) *Server {
	return &Server{
		server: server,
	}
}
