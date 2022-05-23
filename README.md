# Prometheus-BME280-exporter

A prometheus exporter for a BOSH BME280 sensor, this time, written in go :green_heart:

## Wiring the sensor

![image](doc/img/GYBME280_header_pinout.jpg)

VIN, GND, SCL and SDA are the notations on the sensor board of the GYBME280 sensor.

## Running the exporter

After you've successfully mounted the sensor, you have to enable the I2C interface. You can either use `sudo raspi-config` and enable it in the UI or add the following to the `/boot/config.txt

```bash
pi@raspberrypi ~ $ sudo vim /boot/config.txt`

Add to the bottom;

dtparam=i2c_arm=on
dtparam=i2c1=on
```

A reboot of your Pi is required to take changes effect. Run the exporter with:

`docker run -it -v /dev/i2c-1:/dev/i2c-1 --privileged -p 9123:9123 lukasbahr/prometheus-bme280-exporter:74b2a86 --loglevel=info` 

The privileged mode is currently required to grant access to the i2c device. Not the best solution but ad-hoc the first working one.


### Because Kubernetes is noice
For a quick-start to run the exporter in a kubernetes cluster, refer to the [kubernetes/daemonset.yaml](kubernetes/daemonset.yaml). The daemonset is configured to run on the nodes of the cluster with a nodeSelector enabled. If you've installed the sensor on all nodes, leave it empty. 
For a more customizable setup, refer to the helm chart.


## Further reading

- Releaser action used for releasing the charts: [https://github.com/helm/chart-releaser-action](https://github.com/helm/chart-releaser-action)
- Chart releaser used for publishing helm chart: [https://github.com/helm/chart-releaser](https://github.com/helm/chart-releaser)
