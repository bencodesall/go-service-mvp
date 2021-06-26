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

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor


# ==============================================================================
# Development / programming (not engineering) DELETE AFTER MOVING TO ORCHESTRATION
run:
	go run app/app-api/main.go

appbuild:
	cd app/app-api && go build && cd ../../

