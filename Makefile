TAG ?= 1.2.1

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

bundle:
	cd golang-job-type &&\
	racetrack plugin bundle --plugin-version=${TAG} --out=..

deploy-sample:
	racetrack deploy sample-golang-function
