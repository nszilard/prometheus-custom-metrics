package config

func setupDevConfig() {
	options = Options{
		System:      system,
		Environment: Development,
		Port:        8080,
		IsSetup:     true,
	}
}
