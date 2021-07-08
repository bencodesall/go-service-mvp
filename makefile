SHELL := /bin/bash

export PROJECT = bencodesall-starter-kit

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
# Running from within k8s/dev
kind-up:
	kind create cluster --image kindest/node:v1.21.1 --name bencodesall-starter-cluster --config zarf/k8s/dev/kind-config.yaml

kind-down:
	kind delete cluster --name bencodesall-starter-cluster

kind-load:
	kind load docker-image app-api-amd64:1.0 --name bencodesall-starter-cluster
	#kind load docker-image metrics-amd64:1.0 --name bencodesall-starter-cluster

kind-services:
	kustomize build zarf/k8s/dev | kubectl apply -f -

kind-app-api: app-api
	kind load docker-image app-api-amd64:1.0 --name bencodesall-starter-cluster
	kubectl delete pods -lapp=app-api

#kind-metrics: metrics
#	kind load docker-image metrics-amd64:1.0 --name ardan-starter-cluster
#	kubectl delete pods -lapp=sales-api

kind-logs:
	kubectl logs -lapp=app-api --all-containers=true -f

kind-status:
	kubectl get nodes
	kubectl get pods --watch

kind-status-full:
	kubectl describe pod -lapp=app-api

kind-info:
	kubectl cluster-info --context kind-bencodesall-starter-cluster

kind-info-full:
	kubectl cluster-info dump
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

