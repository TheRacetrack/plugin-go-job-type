module github.com/TheRacetrack/plugin-go-job-type/golang-job-type

go 1.20

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/prometheus/client_golang v1.14.0
	github.com/inconshreveable/log15 v0.0.0-20201112154412-8562bdadbbac
	github.com/pkg/errors v0.9.1
	racetrack/job v0.0.0
)

replace racetrack/job => ./handler
