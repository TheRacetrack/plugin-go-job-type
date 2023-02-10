module github.com/TheRacetrack/plugin-go-job-type/golang-job-type

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/inconshreveable/log15 v0.0.0-20201112154412-8562bdadbbac
	github.com/pkg/errors v0.9.1
	racetrack/job v0.0.0
)

replace racetrack/job => ./handler
