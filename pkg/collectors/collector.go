package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
)

// BMECollector defines the struct for the collector that contains pointers
// to prometheus descriptors for each metric you wish to expose.
type BMECollector struct {
	temperature *prometheus.Desc
	humidity    *prometheus.Desc
	pressureMg  *prometheus.Desc
	pressureHPa *prometheus.Desc
	altitude    *prometheus.Desc
}

// NewBMECollector is the constructor for every descriptor and returns a pointer to the collector
func NewBMECollector() *BMECollector {
	return &BMECollector{
		temperature: prometheus.NewDesc("bme280_temperature_celcius",
			"Returns the measured temperature in celsius",
			[]string{"hostname"}, nil,
		),
		humidity: prometheus.NewDesc("bme280_humidity_percent",
			"Returns the measured humidity in percent",
			[]string{"hostname"}, nil,
		),
		pressureMg: prometheus.NewDesc("bme280_pressure_mmHg",
			"Returns the measured and calculated air pressure in mmHg (millimeter of mercury)",
			[]string{"hostname"}, nil,
		),
		pressureHPa: prometheus.NewDesc("bme280_pressure_hpa",
			"Returns the measured and calculated air pressure in hPA (Pascal)",
			[]string{"hostname"}, nil,
		),
		altitude: prometheus.NewDesc("bme280_altitude_meters",
			"Returns shows the measured altitude in meters above sea level (101325 Pa)",
			[]string{"hostname"}, nil,
		),
	}
}

// Describe implements the Describe function of the collector
func (collector *BMECollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.temperature
	ch <- collector.humidity
	ch <- collector.pressureHPa
	ch <- collector.pressureMg
	ch <- collector.altitude
}

// Collect implements required collect function for all prometheus collectors
func (collector *BMECollector) Collect(ch chan<- prometheus.Metric) {
	sensor, intf := InitSensor()
	hostname, _ := GetHostname()
	t := GetSensorTemperature(sensor)
	pHPa := GetSensorPressurePa(sensor)
	pMg := GetSensorPressureMmHg(sensor)
	a := GetSensorAltitude(sensor)
	h := GetSensorHumidityRH(sensor)
	defer intf.Close()
	ch <- prometheus.MustNewConstMetric(collector.temperature, prometheus.GaugeValue, float64(t), hostname)
	ch <- prometheus.MustNewConstMetric(collector.humidity, prometheus.GaugeValue, float64(h), hostname)
	ch <- prometheus.MustNewConstMetric(collector.pressureMg, prometheus.GaugeValue, float64(pMg), hostname)
	ch <- prometheus.MustNewConstMetric(collector.pressureHPa, prometheus.GaugeValue, float64(pHPa), hostname)
	ch <- prometheus.MustNewConstMetric(collector.altitude, prometheus.GaugeValue, float64(a), hostname)
}
