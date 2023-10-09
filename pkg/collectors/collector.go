package collectors

import (
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
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
			[]string{"sensor_id", "hostname"}, nil,
		),
		humidity: prometheus.NewDesc("bme280_humidity_percent",
			"Returns the measured humidity in percent",
			[]string{"sensor_id", "hostname"}, nil,
		),
		pressureMg: prometheus.NewDesc("bme280_pressure_mmHg",
			"Returns the measured and calculated air pressure in mmHg (millimeter of mercury)",
			[]string{"sensor_id", "hostname"}, nil,
		),
		pressureHPa: prometheus.NewDesc("bme280_pressure_hpa",
			"Returns the measured and calculated air pressure in hPA (Pascal)",
			[]string{"sensor_id", "hostname"}, nil,
		),
		altitude: prometheus.NewDesc("bme280_altitude_meters",
			"Returns shows the measured altitude in meters above sea level (101325 Pa)",
			[]string{"sensor_id", "hostname"}, nil,
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
	//temporarily used, more a smell than a feature
	i2cerr := logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	if i2cerr != nil {
		log.Fatal(i2cerr)
	}
	bsbmperr := logger.ChangePackageLogLevel("bsbmp", logger.InfoLevel)
	if bsbmperr != nil {
		log.Fatal(bsbmperr)
	}

	hostname, _ := GetHostname()

	i, err := i2c.NewI2C(0x76, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i.Close()

	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i)
	if err != nil {
		log.Fatal(err)
	}

	id, err := sensor.ReadSensorID()
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("This Bosch Sensortec sensor has signature: 0x%x", id)
	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	pHPa, err := sensor.ReadPressurePa(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	pMg, err := sensor.ReadPressureMmHg(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	a, err := sensor.ReadAltitude(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	supported, h, err := sensor.ReadHumidityRH(bsbmp.ACCURACY_HIGH)
	if supported {
		if err != nil {
			log.Fatal(err)
		}
	}
	ch <- prometheus.MustNewConstMetric(collector.temperature, prometheus.GaugeValue, float64(t), string(id), hostname)
	ch <- prometheus.MustNewConstMetric(collector.humidity, prometheus.GaugeValue, float64(h), string(id), hostname)
	ch <- prometheus.MustNewConstMetric(collector.pressureMg, prometheus.GaugeValue, float64(pMg), string(id), hostname)
	ch <- prometheus.MustNewConstMetric(collector.pressureHPa, prometheus.GaugeValue, float64(pHPa), string(id), hostname)
	ch <- prometheus.MustNewConstMetric(collector.altitude, prometheus.GaugeValue, float64(a), string(id), hostname)
}
