#!/bin/bash

export CGO_ENABLED=0

go build -o "$(pwd)/bin/prom-apcupsd-exporter" -tags netgo -a
