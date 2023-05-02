package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nszilard/prometheus-custom-metrics/internal/pkg"
)

// Normal documentation
// @Summary Responds with a 200 HTTP status code
// @Tags v1
// @Produce text/text
// @Success 200
// @Router /v1/ok [get]
func Normal(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("All good."))
}

// Delay documentation
// @Summary Responds with a 200 HTTP status code but with a random delay
// @Tags v1
// @Produce text/text
// @Success 200
// @Router /v1/delay [get]
func Delay(w http.ResponseWriter, req *http.Request) {
	ms := pkg.GetRandomDuration()

	time.Sleep(ms)
	msg := fmt.Sprintf("Responded after: %v", ms)

	w.Write([]byte(msg))
}

// Exception documentation
// @Summary Responds with a 500 HTTP status code
// @Tags v1
// @Failure 500 {string} string "Oh no, something went wrong!"
// @Router /v1/error [get]
func Exception(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Oh no, something went wrong!"))
}
