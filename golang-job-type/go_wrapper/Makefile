.PHONY: run

run:
	FATMAN_NAME=golang-function FATMAN_VERSION=0.0.1 go run .

perform:
	curl -X POST "http://localhost:7000/pub/fatman/golang-function/latest/api/v1/perform" \
		-H "Content-Type: application/json" \
		-d '{"numbers": [40, 2]}'
