# Makefile as a command wrapper

ARTIFACTS := $(shell find . -name '*.tgz')
GHBASE := "https://lukibahr.github.io/"
REPO := "prometheus-bme280-exporter"
CR := $(shell which cr)

all: login build tag push

include .env
export

package:
	$(CR) package charts

upload:
	$(CR) upload charts	

cleanup:
	rm $(ARTIFACTS)

index:
	$(CR) index --charts-repo $(GHBASE)$(REPO)
