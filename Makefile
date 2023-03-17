TAG ?= 1.3.0

run:
	cd golang-job-type/go_wrapper &&\
	JOB_NAME=golang-function JOB_VERSION=0.0.1 go run .

perform:
	curl -X POST \
		"http://localhost:7000/pub/job/golang-function/latest/api/v1/perform" \
		-H "Content-Type: application/json" \
		-d '{"numbers": [40, 2]}'

build:
	cd golang-job-type &&\
	DOCKER_BUILDKIT=1 docker build \
		-t ghcr.io/theracetrack/racetrack/job-base/golang:latest \
		-f base.Dockerfile .

bundle:
	cd golang-job-type &&\
	racetrack plugin bundle --plugin-version=${TAG} --out=..

deploy-sample:
	racetrack deploy sample-golang-function
