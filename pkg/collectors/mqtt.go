package collectors

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	mqtt_client_id   = "prometheus_bme280_exporter"
	queuing_interval = 60
)

func PubSub(mqtt_broker_host, mqtt_broker_port, mqtt_broker_username, mqtt_broker_password string) {

	// defining callback handlers for messaging
	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		log.Infof("message %s received on topic %s\n", msg.Payload(), msg.Topic())
	}

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		log.Infof("connected on broker %s ", mqtt_broker_host)
	}

	var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		log.Infof("connection to borker %s lost with error: %s\n", mqtt_broker_host, err.Error())
	}

	hostname, _ := GetHostname()

	broker := fmt.Sprintf("tcp://%s:%s", mqtt_broker_host, mqtt_broker_port)
	topicPrefix := fmt.Sprintf("sensors/bme280/%s", hostname)
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(mqtt_client_id)
	options.SetDefaultPublishHandler(messagePubHandler)
	options.SetUsername(mqtt_broker_username)
	options.SetPassword(mqtt_broker_password)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	log.Debugf("initialized mqtt client with host %s port %s username %s password %s ", mqtt_broker_host, mqtt_broker_port, mqtt_broker_username, mqtt_broker_password)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token = client.Subscribe(topicPrefix, 1, nil)
	token.Wait()
	log.Infof("subscribed to topic %s\n", topicPrefix)

	defer client.Disconnect(100)

	// ctx := context.Background()

	sensor := InitSensor()

	go func() {
		for {
			temperature := GetSensorTemperature(sensor)

			token = client.Publish(fmt.Sprintf("%s/temperature", topicPrefix), 0, false, temperature)
			token.Wait()
			log.Infof("sleeping for %d seconds", queuing_interval)
			time.Sleep(queuing_interval * time.Second)

		}
	}()
}
