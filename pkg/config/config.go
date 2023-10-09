package config

type Config struct {
	Port               string
	Accuracy           string
	Loglevel           string
	MQTTBrokerHost     string
	MQTTBrokerPort     string
	MQTTBrokerUsername string
	MQTTBrokerPassword string
}

// New creates a new Config object
func New(port, accuracy, loglevel, mqttbroker, mqttport, mqttusername, mqttpassword string) *Config {
	return &Config{
		Port:               port,
		Accuracy:           accuracy,
		Loglevel:           loglevel,
		MQTTBrokerHost:     mqttbroker,
		MQTTBrokerPort:     mqttport,
		MQTTBrokerUsername: mqttusername,
		MQTTBrokerPassword: mqttpassword,
	}
}
