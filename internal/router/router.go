// Package router provides the REST API router for the service
package router

import (
	chi "github.com/go-chi/chi/v5"

	"github.com/nszilard/prometheus-custom-metrics/internal/handlers"
	"github.com/nszilard/prometheus-custom-metrics/internal/router/swagger"
	v1 "github.com/nszilard/prometheus-custom-metrics/internal/router/v1"
)

// Create will instantiate the REST API router for this service
func Create() chi.Router {
	r := chi.NewRouter()

	// Handle common routes
	r.Get("/alive", handlers.Alive)
	r.Get("/ready", handlers.Ready)
	r.Get("/metrics", handlers.Metrics)

	// Mount the swagger documentation
	r.Mount("/swagger", swagger.Attach())

	// Mount REST endpoints
	r.Mount("/v1", v1.Attach())

	return r
}
