package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

 type client struct {
	i2c *i2c.I2C
	//sensor *bsbmp.SensorBME280
	sensor *bsbmp.BMP
 }

var temperatureGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "bme280",
		Name:      "temperature",
		Help:      "Temperature measured in celcius from bme280 sensor",
	},
)

var humidityGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "bme280",
		Name:      "humidity",
		Help:      "Humidity measured in percent from bme280 sensor",
	},
)

type customCollector struct {
}

func (fc customCollector) Collect(c chan<- prometheus.Metric) {

	f, err := client
	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_HIGH)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("measured temperature = %v*C", t)
	temperatureGauge.Set(t)
	c <- temperatureGauge

	l, err := fritzClient.List()
	if err != nil {
		log.Println("Unable to collect data:", err)
		return
	}
	temperatureGauge.Set(tc)
	c <- temperatureGauge
}

func (fc customCollector) Describe(c chan<- *prometheus.Desc) {
	temperatureGauge.Describe(c)
}

func main() {
	fc := customCollector{}
	i, err := i2c.NewI2C(0x76, 1)
	if err != nil {
		log.Fatal(err)
	}
	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i)
	if err != nil {
		log.Fatal(err)
	}
	id, err := sensor.ReadSensorID()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("This Bosch Sensortec sensor has signature: 0x%x", id)
	c := client{
		i2c:    i,
		sensor: sensor,
	}
	defer i.Close()
	go func() {
		for {
			time.Sleep(10 * time.Minute)
		}
	}()

	err := prometheus.Register(fc)
	if err != nil {
		log.Fatalln(err)
	}
	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":9103", nil); err != nil {
		log.Fatalln(err)
	}
}

