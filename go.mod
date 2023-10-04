module github.com/lukibahr/Prometheus-BME280-exporter

go 1.20.2

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/d2r2/go-bsbmp v0.0.0-20190515110334-3b4b3aea8375
	github.com/d2r2/go-i2c v0.0.0-20191123181816-73a8a799d6bc
	//github.com/d2r2/go-logger v0.0.0-20181221090742-9998a510495e
	github.com/d2r2/go-logger v0.0.0-20210606094344-60e9d1233e22

	github.com/prometheus/client_golang v1.12.2
	github.com/shiena/ansicolor v0.0.0-20230509054315-a9deabde6e02
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.7.0
)
