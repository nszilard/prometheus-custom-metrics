package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Register will ensure that the custom metrics are registered with Prometheus
func Register() {
	// Counters
	prometheus.DefaultRegisterer.MustRegister(ApplicationError)
	prometheus.DefaultRegisterer.MustRegister(EndpointAccessed)

	// Gauges
	prometheus.DefaultRegisterer.MustRegister(RequestsInFlight)

	// Histograms
	prometheus.DefaultRegisterer.MustRegister(ResponseDuration)
}
