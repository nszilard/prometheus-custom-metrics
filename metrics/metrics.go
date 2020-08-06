package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// ApplicationError metrics counts the number of errors the system has experienced
var ApplicationError = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "application_error",
		Help: "Number of errors observed in the system.",
	},
	[]string{"system", "endpoint", "code"},
)

// IncrementApplicationError increments the appropriate prometheus counter metric
func IncrementApplicationError(system, endpoint string, code int) {
	ApplicationError.With(prometheus.Labels{
		"system":   system,
		"endpoint": endpoint,
		"code":     strconv.Itoa(code),
	}).Inc()
}

// EndpointAccessed metrics counts how many times an endpoint was hit
var EndpointAccessed = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "endpoint_accessed",
		Help: "Number of times an endpoint was accessed.",
	},
	[]string{"system", "endpoint"},
)

// IncrementEndpointAccessed increments the appropriate prometheus counter metric
func IncrementEndpointAccessed(system, endpoint string) {
	EndpointAccessed.With(prometheus.Labels{
		"system":   system,
		"endpoint": endpoint,
	}).Inc()
}

// ActiveDatabaseConnection stores the active database connections at any given time
var ActiveDatabaseConnection = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "active_database_connection",
		Help: "Number of active database connections currently in the system.",
	},
	[]string{"system"},
)

// IncrementActiveDatabaseConnection increments the appropriate prometheus gauge metric
func IncrementActiveDatabaseConnection(system string) {
	ActiveDatabaseConnection.With(prometheus.Labels{
		"system": system,
	}).Inc()
}

// DecrementActiveDatabaseConnection decrements the appropriate prometheus gauge metric
func DecrementActiveDatabaseConnection(system string) {
	ActiveDatabaseConnection.With(prometheus.Labels{
		"system": system,
	}).Dec()
}

// ResponseDuration stores the observed response durations in seconds
var ResponseDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "response_duration_seconds",
		Help:    "Response distribution for the system in seconds.",
		Buckets: prometheus.ExponentialBuckets(0.01, 4, 7),
	},
	[]string{"system", "endpoint"},
)

// ObserveResponseDuration stores the duration in the appropriate bucket for the prometheus histogram metric
func ObserveResponseDuration(system, endpoint string, duration float64) {
	ResponseDuration.With(prometheus.Labels{
		"system":   system,
		"endpoint": endpoint,
	}).Observe(duration)
}
