package metrics

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/nszilard/prometheus-custom-metrics/config"
	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.handleGet(w, req)
	}
}

func (h *handler) handleGet(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case home:
		h.handleCounterEndpoint(w, req, home)
	case metrics:
		h.handleMetrics(w, req)
	case errorNotFound:
		h.handleError(w, req, errorNotFound)
	case add:
		h.handleGaugeEndpointInc(w, req, add)
	case sub:
		h.handleGaugeEndpointDec(w, req, sub)
	case responseDuration:
		h.handleHistogramEndpoint(w, req, responseDuration)
	}
}

func (h *handler) handleMetrics(w http.ResponseWriter, req *http.Request) {
	promhttp.Handler().ServeHTTP(w, req)
}

func (h *handler) handleCounterEndpoint(w http.ResponseWriter, req *http.Request, endpoint string) {
	IncrementEndpointAccessed(config.System, endpoint)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) handleError(w http.ResponseWriter, req *http.Request, endpoint string) {
	IncrementApplicationError(config.System, endpoint, http.StatusNotFound)
	w.WriteHeader(http.StatusNotFound)
}

func (h *handler) handleGaugeEndpointInc(w http.ResponseWriter, req *http.Request, endpoint string) {
	IncrementActiveDatabaseConnection(config.System)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) handleGaugeEndpointDec(w http.ResponseWriter, req *http.Request, endpoint string) {
	DecrementActiveDatabaseConnection(config.System)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) handleHistogramEndpoint(w http.ResponseWriter, req *http.Request, endpoint string) {
	start := time.Now()

	func() {
		rand.Seed(start.UnixNano())
		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}()

	duration := time.Since(start)
	ObserveResponseDuration(config.System, endpoint, duration.Seconds())
	w.WriteHeader(http.StatusOK)
}

// listen will start a test server listening to a local port.
func getTestServer() (*httptest.Server, error) {
	l, err := net.Listen("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, err
	}

	ts := httptest.NewUnstartedServer(&handler{})
	ts.Listener = l
	return ts, nil
}
