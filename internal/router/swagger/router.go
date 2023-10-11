package swagger

import (
	chi "github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/nszilard/prometheus-custom-metrics/docs" // docs is generated by Swag CLI and has to be imported
)

// Attach handles adding the routes for this endpoint
func Attach() chi.Router {
	r := chi.NewRouter()

	r.Get("/*", httpSwagger.WrapHandler)

	return r
}