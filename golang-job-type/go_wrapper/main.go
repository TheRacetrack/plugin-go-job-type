package main

import (
	handler "racetrack/fatman"
)

func main() {
	err := WrapAndServe(handler.Perform)
	if err != nil {
		panic(err)
	}
}
