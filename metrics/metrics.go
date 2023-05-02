package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// -----------------------------------
// Counters
// -----------------------------------

// ApplicationError metrics counts the number of errors the system has experienced
var ApplicationError = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "application_error",
		Help: "Number of errors observed in the system.",
	},
	[]string{labelSystem, labelEndpoint, labelMethod, labelCode},
)

// IncrementApplicationError increments the appropriate prometheus counter metric
func IncrementApplicationError(system, endpoint, method string, code int) {
	ApplicationError.With(prometheus.Labels{
		labelSystem:   system,
		labelEndpoint: endpoint,
		labelMethod:   method,
		labelCode:     strconv.Itoa(code),
	}).Inc()
}

// EndpointAccessed metrics counts how many times an endpoint was hit
var EndpointAccessed = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "endpoint_accessed",
		Help: "Number of times an endpoint was accessed.",
	},
	[]string{labelSystem, labelEndpoint, labelMethod},
)

// IncrementEndpointAccessed increments the appropriate prometheus counter metric
func IncrementEndpointAccessed(system, endpoint, method string) {
	EndpointAccessed.With(prometheus.Labels{
		labelSystem:   system,
		labelEndpoint: endpoint,
		labelMethod:   method,
	}).Inc()
}

// -----------------------------------
// Gauges
// -----------------------------------

// RequestsInFlight stores the number of concurrent in-flight requests
var RequestsInFlight = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "requests_inflight",
		Help: "The number of inflight requests being handled at the same time.",
	},
	[]string{labelSystem},
)

// IncrementInflightRequests increments the appropriate prometheus gauge metric
func IncrementInflightRequests(system string) {
	RequestsInFlight.With(prometheus.Labels{
		labelSystem: system,
	}).Inc()
}

// DecrementInflightRequests decrements the appropriate prometheus gauge metric
func DecrementInflightRequests(system string) {
	RequestsInFlight.With(prometheus.Labels{
		labelSystem: system,
	}).Dec()
}

// -----------------------------------
// Histograms
// -----------------------------------

// ResponseDuration stores the observed response durations in seconds
var ResponseDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "response_duration_seconds",
		Help:    "Response duration distribution in seconds.",
		Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 1.5, 2},
	},
	[]string{labelSystem, labelEndpoint},
)

// ObserveResponseDuration stores the duration in the appropriate bucket for the prometheus histogram metric
func ObserveResponseDuration(system, endpoint string, duration float64) {
	ResponseDuration.With(prometheus.Labels{
		labelSystem:   system,
		labelEndpoint: endpoint,
	}).Observe(duration)
}
