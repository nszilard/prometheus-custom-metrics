package handlers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics documentation
// @Summary Prometheus Metrics
// @Description Metrics is an http.Handler instance to expose Prometheus metrics via HTTP.
// @ID metrics
// @Tags Common
// @Success 200
// @Router /metrics [get]
func Metrics(w http.ResponseWriter, req *http.Request) {
	promhttp.Handler().ServeHTTP(w, req)
}
