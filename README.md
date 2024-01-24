# Racetrack Plugin: Go Job Type

This is a plugin for [Racetrack](https://github.com/TheRacetrack/racetrack)
which extends it with Golang Job Type.
It's a language wrapper converting your code written in Go to a Job web service.

## Setup
1. Install `racetrack` client and generate ZIP plugin by running `make bundle`.

2. Activate the plugin in Racetrack Dashboard Admin page
  by uploading the zipped plugin file.

## Usage
You can deploy sample Go job by running:
```bash
racetrack deploy sample
```

See [how to use Go job type](./docs/job_golang.md).
