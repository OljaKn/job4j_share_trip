package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"job4j.ru/go-share-trip/internal/observability/metrics"
	"job4j.ru/go-share-trip/internal/service"
)

type Server struct {
	server   *service.Service
	registry *prometheus.Registry
	metrics  *metrics.Metrics
}

func NewServer(server *service.Service, registry *prometheus.Registry, metrics *metrics.Metrics) *Server {
	return &Server{
		server:   server,
		registry: registry,
		metrics:  metrics,
	}
}
