# Racetrack Plugin: Go Job Type

This is a plugin for Racetrack which extends it with Golang Job Type.
It's a language wrapper converting your code written in Go to a Fatman web service.

## Setup
1. Make sure that current version of language wrapper docker image
  (provided by plugin) is pushed to your Docker registry,
  which is used by your Racetrack instance. 
  Do it by pushing to public registry: `make push`  
  or if you want to use private registry, run `make env-template`,
  fill in `.env` file and run `make push-private-registry`

2. Activate the plugin in Racetrack, 
  add the following to your image-builder configuration (in kustomize ConfigMap):

```yaml
plugins:
- name: golang-job-type
  git_remote: https://github.com/TheRacetrack/plugin-go-job-type
  git_ref: '1.0.0'
  git_directory: golang-job-type
```

## Usage
You can deploy sample Go job by running:
```bash
racetrack deploy sample-golang-function <RACETRACK_URL>
```
