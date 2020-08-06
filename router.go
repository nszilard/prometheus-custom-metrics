package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nszilard/prometheus-custom-metrics/config"
	"github.com/nszilard/prometheus-custom-metrics/metrics"
	"github.com/nszilard/prometheus-custom-metrics/random"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// requestHandler takes a request and delegates the processing to the handler
type requestHandler struct{}

func (h *requestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.handleGet(w, req)
	}
}

func (h *requestHandler) handleGet(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case pathHome:
		h.getHome(w, req, pathHome)
	case pathHealthcheck:
		h.getHealthCheck(w, req, pathHealthcheck)
	case pathVersion:
		h.getVersion(w, req, pathVersion)
	case pathRandom:
		h.getRandom(w, req, pathRandom)
	case pathError:
		h.getError(w, req, pathError)
	case pathCreateConnection:
		h.getConnectionAdd(w, req, pathCreateConnection)
	case pathTerminateConnection:
		h.getConnectionSub(w, req, pathTerminateConnection)
	case pathMetrics:
		h.getMetrics(w, req)
	}
}

func (*requestHandler) getMetrics(w http.ResponseWriter, req *http.Request) {
	promhttp.Handler().ServeHTTP(w, req)
}

func (*requestHandler) getHome(w http.ResponseWriter, req *http.Request, endpoint string) {
	start := time.Now()
	metrics.IncrementEndpointAccessed(config.System, endpoint)

	w.Write([]byte("Sample app to showcase Prometheus custom metrics"))
	metrics.ObserveResponseDuration(config.System, endpoint, time.Since(start).Seconds())
}

func (*requestHandler) getHealthCheck(w http.ResponseWriter, req *http.Request, endpoint string) {
	metrics.IncrementEndpointAccessed(config.System, endpoint)
	w.Write([]byte("200 OK"))
}

func (*requestHandler) getVersion(w http.ResponseWriter, req *http.Request, endpoint string) {
	metrics.IncrementEndpointAccessed(config.System, endpoint)
	w.Write([]byte(fmt.Sprintf("%s: v%s", config.System, VERSION)))
}

func (*requestHandler) getConnectionAdd(w http.ResponseWriter, req *http.Request, endpoint string) {
	metrics.IncrementEndpointAccessed(config.System, endpoint)
	metrics.IncrementActiveDatabaseConnection(config.System)
	w.Write([]byte("Increasing database connection by 1"))
}

func (*requestHandler) getConnectionSub(w http.ResponseWriter, req *http.Request, endpoint string) {
	metrics.IncrementEndpointAccessed(config.System, endpoint)
	metrics.DecrementActiveDatabaseConnection(config.System)
	w.Write([]byte("Decreasing database connection by 1"))
}

func (*requestHandler) getError(w http.ResponseWriter, req *http.Request, endpoint string) {
	metrics.IncrementEndpointAccessed(config.System, endpoint)

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Oh no, something went wrong"))

	metrics.IncrementApplicationError(config.System, endpoint, http.StatusNotFound)
}

func (*requestHandler) getRandom(w http.ResponseWriter, req *http.Request, endpoint string) {
	start := time.Now()
	metrics.IncrementEndpointAccessed(config.System, endpoint)

	out, err := random.Generate()
	if err != nil {
		metrics.IncrementApplicationError(config.System, endpoint, http.StatusInternalServerError)
		metrics.ObserveResponseDuration(config.System, endpoint, time.Since(start).Seconds())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(out)
	metrics.ObserveResponseDuration(config.System, endpoint, time.Since(start).Seconds())
}
