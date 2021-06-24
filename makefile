SHELL := /bin/bash

run:
	go run app/application/main.go

tidy:
	go mod tidy
	go mod vendor

appbuild:
	cd app/application && go build && cd ../../

