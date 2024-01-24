package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Service                    string `json:"service"`
	JobName                    string `json:"job_name"`
	JobVersion                 string `json:"job_version"`
	GitVersion                 string `json:"git_version"`
	DeployedByRacetrackVersion string `json:"deployed_by_racetrack_version"`
	Status                     string `json:"status"`
	DeploymentTimestamp        int    `json:"deployment_timestamp"`
}

type LiveResponse struct {
	Status              string `json:"status"`
	DeploymentTimestamp int    `json:"deployment_timestamp"`
}

type ReadyResponse struct {
	Status string `json:"status"`
}

func HealthHandler(c *gin.Context) {
	deploymentTimestamp, _ := strconv.Atoi(os.Getenv("JOB_DEPLOYMENT_TIMESTAMP"))
	output := &HealthResponse{
		Service:                    "job",
		JobName:                    os.Getenv("JOB_NAME"),
		JobVersion:                 os.Getenv("JOB_VERSION"),
		GitVersion:                 os.Getenv("GIT_VERSION"),
		DeployedByRacetrackVersion: os.Getenv("DEPLOYED_BY_RACETRACK_VERSION"),
		DeploymentTimestamp:        deploymentTimestamp,
		Status:                     "pass",
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(output)
}

func LiveHandler(c *gin.Context) {
	deploymentTimestamp, _ := strconv.Atoi(os.Getenv("JOB_DEPLOYMENT_TIMESTAMP"))
	output := &LiveResponse{
		Status:              "live",
		DeploymentTimestamp: deploymentTimestamp,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(output)
}

func ReadyHandler(c *gin.Context) {
	output := &ReadyResponse{
		Status: "ready",
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(output)
}
