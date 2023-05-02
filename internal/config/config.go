package config

import (
	"os"
	"sync"

	"github.com/nszilard/log"
)

const (
	system = "prometheus-custom-metrics"

	// Development defines the dev environment name
	Development = "development"

	// Production defines the prod environment name
	Production = "production"
)

var (
	// options holds the config we configured for the environment in question.
	options Options

	// mux is a mutex making sure we're not trying to set up the config simultaneously from several locations.
	mux = sync.Mutex{}
)

// Get returns the configuration for the applicable environment.
func Get() Options {
	mux.Lock()
	defer mux.Unlock()

	if !options.IsSetup {
		env := os.Getenv("APPLICATION_ENV")

		switch env {
		case Development:
			setupDevConfig()
		case Production:
			setupProdConfig()
		default:
			log.Panicf("%v: no env: %q", system, env)
			panic("config: no environment set")
		}

		log.Infof("%v: config has been initialized for the environment: %q", system, env)
	}

	return options
}
