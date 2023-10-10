package collectors

import (
	"fmt"
	"os"

	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
	log "github.com/sirupsen/logrus"
)

func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return hostname, nil
}

func InitSensor() (*bsbmp.BMP, *i2c.I2C) {
	// temporarily used, more a smell than a feature
	i2cerr := logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	if i2cerr != nil {
		log.Fatal(i2cerr)
	}
	bsbmperr := logger.ChangePackageLogLevel("bsbmp", logger.InfoLevel)
	if bsbmperr != nil {
		log.Fatal(bsbmperr)
	}

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
	log.Debugf("This Bosch Sensortec sensor has signature: 0x%x", id)
	return sensor, i
}

func GetSensorTemperature(sensor *bsbmp.BMP) float32 {
	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
func GetSensorPressureMmHg(sensor *bsbmp.BMP) float32 {
	pMg, err := sensor.ReadPressureMmHg(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	return pMg
}
func GetSensorPressurePa(sensor *bsbmp.BMP) float32 {
	pHPa, err := sensor.ReadPressurePa(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	return pHPa
}
func GetSensorAltitude(sensor *bsbmp.BMP) float32 {
	a, err := sensor.ReadAltitude(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func GetSensorHumidityRH(sensor *bsbmp.BMP) float32 {
	supported, h, err := sensor.ReadHumidityRH(bsbmp.ACCURACY_HIGH)
	if supported {
		if err != nil {
			log.Fatal(err)
		}
	}
	return h
}
