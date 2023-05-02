package middlewares

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/nszilard/prometheus-custom-metrics/metrics"
)

// ------------------------------------------------------------------------------------
// Metrics is a middleware that can collect application metrics from the service.
// It requires a system string to identify which service it is collecting mertics for.
// ------------------------------------------------------------------------------------

// Metrics middleware takes care about invoking the functions to collect our application metrics
func Metrics(system string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Keep the time
			start := time.Now()

			// Wrap the response writer
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			// Keep the number of times the endpoint was accessed
			metrics.IncrementEndpointAccessed(system, r.URL.Path, r.Method)

			// Keep track the number of requests the system handles at any given moment
			metrics.IncrementInflightRequests(system)

			defer func() {
				metrics.DecrementInflightRequests(system)

				// Keep the time it took to respond to the request
				metrics.ObserveResponseDuration(system, r.URL.Path, time.Since(start).Seconds())

				// Keep the error response code
				if code := filterErrorResponseCode(ww.Status()); code != 0 {
					metrics.IncrementApplicationError(system, r.URL.Path, r.Method, code)
				}
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}

// filterErrorResponseCode returns the erroneous HTTP status codes as registered with IANA.
func filterErrorResponseCode(code int) int {
	switch code {
	// Response code: 1xx
	case
		http.StatusContinue,
		http.StatusSwitchingProtocols,
		http.StatusProcessing,
		http.StatusEarlyHints:
		return 0

	// Response code: 2xx
	case
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed:
		return 0

	// Response code: 3xx
	case
		http.StatusMultipleChoices,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusNotModified,
		http.StatusUseProxy,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect:
		return 0

	// Anything else counts as an error code
	default:
		return code
	}
}
