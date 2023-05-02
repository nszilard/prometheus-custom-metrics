package main

import (
	"fmt"
	"net/http"

	"github.com/nszilard/log"

	"github.com/nszilard/prometheus-custom-metrics/internal/config"
	"github.com/nszilard/prometheus-custom-metrics/internal/router"
	"github.com/nszilard/prometheus-custom-metrics/metrics"
)

func init() {
	// Register Prometheus metrics
	metrics.Register()
}

// -----------------------------------
// Swagger annotations
// -----------------------------------
// @title Prometheus Custom Metrics
// @version 1.0
// @description API documentation for the 'Prometheus Custom Metrics' application.
// @BasePath /

func main() {
	conf := config.Get()
	port := fmt.Sprintf(":%v", conf.Port)

	router := router.Create()
	log.Infof("%v: listening on port %v", conf.System, port[1:])

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("%v: error while serving api: %v", conf.System, err)
		panic(err)
	}
}
