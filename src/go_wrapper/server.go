package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"github.com/pkg/errors"
)

// WrapAndServe embeds given function in a HTTP server and listens for requests
func WrapAndServe(entrypoint EntrypointHandler) error {
	performHandler := buildHandler(entrypoint)

	jobName := os.Getenv("JOB_NAME")
	// Serve endpoints at raw path (to facilitate debugging) and prefixed path (when accessed through PUB).
	// Accept any version so that job can be called by its many version names ("latest", "1.x").
	baseUrls := []string{
		fmt.Sprintf("/pub/job/%s/:version", jobName),
		"",
	}

	gin.SetMode(gin.ReleaseMode) //Hide debug routings
	router := gin.New()
	router.Use(gin.Recovery())

	for _, baseUrl := range baseUrls {

		router.POST(baseUrl+"/api/v1/perform", performHandler)
		router.GET(baseUrl+"/health", HealthHandler)
		router.GET(baseUrl+"/live", LiveHandler)
		router.GET(baseUrl+"/ready", ReadyHandler)
		router.GET(baseUrl+"/metrics", MetricHandler())
		MountOpenApi(router, baseUrl)
	}

	router.Use(gin.Logger())

	listenAddress := "0.0.0.0:7000"
	log.Info("Listening on", log.Ctx{
		"listenAddress": listenAddress,
		"baseUrls":      baseUrls,
	})
	if err := router.Run(listenAddress); err != nil {
		log.Error("Serving http", log.Ctx{"error": err})
		return errors.Wrap(err, "Failed to serve")
	}
	return nil
}

type EntrypointHandler func(input map[string]interface{}) (interface{}, error)

func buildHandler(entrypointHandler EntrypointHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Debug("Perform request received")

		var input map[string]interface{}
		err := json.NewDecoder(c.Request.Body).Decode(&input)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := entrypointHandler(input)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(c.Writer).Encode(output)
	}
}

func wrapHandler(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
