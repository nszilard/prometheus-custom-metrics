package config

// Options is the main config type that holds configuration options for the service.
type Options struct {
	System      string
	Environment string
	Port        int
	IsSetup     bool
}
