run:
	cd src &&\
	JOB_NAME=golang-function JOB_VERSION=0.0.1 go run .

perform:
	curl -X POST \
		"http://localhost:7000/pub/job/golang-function/latest/api/v1/perform" \
		-H "Content-Type: application/json" \
		-d '{"numbers": [40, 2]}'

bundle:
	cd src &&\
	racetrack plugin bundle --plugin-version=${TAG} --out=..

deploy-sample:
	racetrack deploy sample
