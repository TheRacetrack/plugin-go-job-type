package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type HealthResponse struct {
	Service                    string `json:"service"`
	FatmanName                 string `json:"fatman_name"`
	FatmanVersion              string `json:"fatman_version"`
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

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	deploymentTimestamp, _ := strconv.Atoi(os.Getenv("FATMAN_DEPLOYMENT_TIMESTAMP"))
	output := &HealthResponse{
		Service:                    "fatman",
		FatmanName:                 os.Getenv("FATMAN_NAME"),
		FatmanVersion:              os.Getenv("FATMAN_VERSION"),
		GitVersion:                 os.Getenv("GIT_VERSION"),
		DeployedByRacetrackVersion: os.Getenv("DEPLOYED_BY_RACETRACK_VERSION"),
		DeploymentTimestamp:        deploymentTimestamp,
		Status:                     "pass",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func LiveHandler(w http.ResponseWriter, req *http.Request) {
	deploymentTimestamp, _ := strconv.Atoi(os.Getenv("FATMAN_DEPLOYMENT_TIMESTAMP"))
	output := &LiveResponse{
		Status:              "live",
		DeploymentTimestamp: deploymentTimestamp,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func ReadyHandler(w http.ResponseWriter, req *http.Request) {
	output := &ReadyResponse{
		Status: "ready",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
