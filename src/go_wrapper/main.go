package main

import (
	handler "racetrack/job"
)

func main() {
	err := WrapAndServe(handler.Perform)
	if err != nil {
		panic(err)
	}
}
