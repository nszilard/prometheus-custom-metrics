package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nszilard/prometheus-custom-metrics/config"
	"github.com/nszilard/prometheus-custom-metrics/metrics"
)

func init() {
	metrics.Register()
}

func main() {
	port := fmt.Sprintf(":%v", config.GetConfig().Port)
	log.Printf("%s: listening on port %v", config.System, port[1:])

	log.Fatal(http.ListenAndServe(port, &requestHandler{}))
}
