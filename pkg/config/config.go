package config

type Config struct {
	Port     string
	Accuracy string
	Loglevel string
}

// New creates a new Config object
func New(port string, accuracy, loglevel string) *Config {
	return &Config{
		Port:     port,
		Accuracy: accuracy,
		Loglevel: loglevel,
	}
}
