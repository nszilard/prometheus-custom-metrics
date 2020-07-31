package config

import (
	"os"
	"sync"
)

const (
	// System is the current system
	System = "prometheus-metrics-example"

	// Testing defines our testing environment.
	Testing = "testing"

	// Production defines our production environment.
	Production = "production"

	// TestPort defines the port to use in unit testing
	TestPort = "9999"
)

// Configuration is our main configuration type, holding internal configurations.
type Configuration struct {
	Environment string
	Port        int
	IsSetup     bool
}

// c holds the config we configured for the environment in question.
var conf Configuration

// mux is a mutex making sure we're not trying to set up the config simultaneously from several locations.
var mux = sync.Mutex{}

// GetConfig will return a C object representing the configuration for the current system environment.
func GetConfig() Configuration {
	mux.Lock()
	defer mux.Unlock()

	if !conf.IsSetup {
		switch os.Getenv("APPLICATION_ENV") {
		case Production:
			setupProdConfig()
		default:
			setupTestingConfig()
		}
	}

	return conf
}

func setupTestingConfig() {
	conf = Configuration{
		Environment: Testing,
		Port:        8500,
		IsSetup:     true,
	}
}

func setupProdConfig() {
	conf = Configuration{
		Environment: Production,
		Port:        8080,
		IsSetup:     true,
	}
}
