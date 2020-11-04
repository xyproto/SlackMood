#!/bin/sh
go build -mod=vendor && ./server -c config/custom.yml -b :8000
