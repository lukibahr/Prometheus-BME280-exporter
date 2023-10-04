package config

type Config struct {
	Port        string
	Accuracy    string
	Loglevel    string
	Environment string
	Location    string
}

// New creates a new Config object
func New(port, accuracy, loglevel, environment, location string) *Config {
	return &Config{
		Port:        port,
		Accuracy:    accuracy,
		Loglevel:    loglevel,
		Environment: environment,
		Location:    location,
	}
}
