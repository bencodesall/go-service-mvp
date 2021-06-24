SHELL := /bin/bash

run:
	go run app/app-api/main.go

tidy:
	go mod tidy
	go mod vendor

appbuild:
	cd app/app-api && go build && cd ../../

