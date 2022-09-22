TAG ?= 1.1.0

-include .env

run:
	cd golang-job-type/go_wrapper &&\
	FATMAN_NAME=golang-function FATMAN_VERSION=0.0.1 go run .

perform:
	curl -X POST \
		"http://localhost:7000/pub/fatman/golang-function/latest/api/v1/perform" \
		-H "Content-Type: application/json" \
		-d '{"numbers": [40, 2]}'

build:
	cd golang-job-type &&\
	DOCKER_BUILDKIT=1 docker build \
		-t ghcr.io/theracetrack/racetrack/fatman-base/golang:latest \
		-f base.Dockerfile .

push: build
	docker login ghcr.io
	docker tag ghcr.io/theracetrack/racetrack/fatman-base/golang:latest ghcr.io/theracetrack/racetrack/fatman-base/golang:$(TAG)
	docker push ghcr.io/theracetrack/racetrack/fatman-base/golang:$(TAG)

push-local: build
	docker tag ghcr.io/theracetrack/racetrack/fatman-base/golang:latest localhost:5000/racetrack/fatman-base/golang:$(TAG)
	docker push localhost:5000/racetrack/fatman-base/golang:$(TAG)

push-private-registry: build
	docker login ${REGISTRY}
	docker tag ghcr.io/theracetrack/racetrack/fatman-base/golang:latest ${REGISTRY}/fatman-base/golang:$(TAG)
	docker push ${REGISTRY}/fatman-base/golang:$(TAG)

push-all: push push-local push-private-registry

env-template:
	cp -n .env.dist .env
	@echo "Now fill in the .env file with your settings"

bundle:
	cd golang-job-type &&\
	racetrack-plugin-bundler bundle

deploy-sample:
	racetrack deploy sample-golang-function docker
