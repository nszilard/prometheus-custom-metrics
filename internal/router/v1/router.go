package v1

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/nszilard/prometheus-custom-metrics/internal/config"
	"github.com/nszilard/prometheus-custom-metrics/internal/handlers/v1"
	"github.com/nszilard/prometheus-custom-metrics/middlewares"
)

// Attach handles adding the routes for this endpoint
func Attach() chi.Router {
	r := chi.NewRouter()

	// Set up middlewares for the router
	r.Use(middleware.Logger)
	r.Use(middlewares.Metrics(config.Get().System))

	// GET Endpoints
	r.Get("/ok", handlers.Normal)
	r.Get("/delay", handlers.Delay)
	r.Get("/error", handlers.Exception)

	return r
}
