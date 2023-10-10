# syntax=docker/dockerfile:1
FROM gcr.io/distroless/base

LABEL maintainer="hello@lukasbahr.de"

COPY prometheus-bme280-exporter /usr/bin/prometheus-bme280-exporter
ENTRYPOINT ["/usr/bin/prometheus-bme280-exporter"]