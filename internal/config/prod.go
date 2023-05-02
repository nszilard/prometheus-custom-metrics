package config

func setupProdConfig() {
	options = Options{
		System:      system,
		Environment: Production,
		Port:        8080,
		IsSetup:     true,
	}
}
