package handlers

import "net/http"

// Alive documentation
// @Summary Kubernetes Alive probe
// @Description Responds to the Kubernetes alive requests
// @ID alive
// @Tags Common
// @Produce text/text
// @Success 200 {string} string "OK"
// @Router /alive [get]
func Alive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// Ready documentation
// @Summary Kubernetes Ready probe
// @Description Responds to the Kubernetes ready requests
// @ID ready
// @Tags Common
// @Produce text/text
// @Success 200 {string} string "OK"
// @Router /ready [get]
func Ready(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
