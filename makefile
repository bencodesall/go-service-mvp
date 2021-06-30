SHELL := /bin/bash

export PROJECT = ardan-starter-kit

# ==============================================================================
# Testing running system

# // To generate a private/public key PEM file.
# openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# openssl rsa -pubout -in private.pem -out public.pem
# ./sales-admin genkey

# curl --user "admin@example.com:gophers" http://localhost:3000/v1/users/token/54bb2165-71e1-41a6-af3e-7da4a0e1e2c1
# export TOKEN="COPY TOKEN STRING FROM LAST CALL"
# curl -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/v1/users

# hey -m GET -c 100 -n 10000 -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/v1/users
# zipkin: http://localhost:9411
# expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,mem:memstats.Alloc"

# ==============================================================================
# Building containers

# $(shell git rev-parse --short HEAD)
VERSION := 1.0
VREF := $(shell git rev-parse --short HEAD)

all: app-api

app-api:
	docker build \
		-f zarf/docker/dockerfile.app-api \
		-t app-api-amd64:$(VERSION) \
		--build-arg VCS_REF=$(VREF) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Running tests locally
test:
	go test -v ./... -count=1
	staticcheck ./...

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor


# ==============================================================================
# Development / programming (not engineering) DELETE AFTER MOVING TO ORCHESTRATION
run:
	go run app/app-api/main.go

runadmin:
	go run app/admin/main.go

buildapi:
	go build app/app-api/main.go

