package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (h *Server) Route(route fiber.Router) {
	route.Get("/ready", h.Ready)
	route.Post("/trip/create", h.CreateTrip)
	route.Get("/trip/:id", h.GetTrip)
	route.Post("/trip/publish", h.PublishTrip)
	route.Get("/metrics", adaptor.HTTPHandler(promhttp.HandlerFor(h.registry, promhttp.HandlerOpts{})))
}
