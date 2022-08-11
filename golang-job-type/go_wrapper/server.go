package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/inconshreveable/log15"
	"github.com/pkg/errors"
)

// WrapAndServe embeds given function in a HTTP server and listens for requests
func WrapAndServe(entrypoint EntrypointHandler) error {
	performHandler := buildHandler(entrypoint)

	fatmanName := os.Getenv("FATMAN_NAME")
	// Serve endpoints at raw path (to facilitate debugging) and prefixed path (when accessed through PUB).
	// Accept any version so that fatman can be called by its many version names ("latest", "1.x").
	baseUrls := []string{
		fmt.Sprintf("/pub/fatman/%s/{version}", fatmanName),
		"",
	}

	router := mux.NewRouter()

	for _, baseUrl := range baseUrls {

		router.HandleFunc(baseUrl+"/api/v1/perform", performHandler)
		router.HandleFunc(baseUrl+"/health", HealthHandler)
		router.HandleFunc(baseUrl+"/live", LiveHandler)
		router.HandleFunc(baseUrl+"/ready", ReadyHandler)
		MountOpenApi(router, baseUrl)
	}

	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("Request", log.Ctx{
				"method": r.Method,
				"uri":    r.RequestURI,
				"ip":     r.RemoteAddr,
			})
			next.ServeHTTP(w, r)
		})
	}
	router.Use(loggingMiddleware)

	listenAddress := "0.0.0.0:7000"
	log.Info("Listening on", log.Ctx{
		"listenAddress": listenAddress,
		"baseUrls":      baseUrls,
	})
	if err := http.ListenAndServe(listenAddress, router); err != nil {
		log.Error("Serving http", log.Ctx{"error": err})
		return errors.Wrap(err, "Failed to serve")
	}
	return nil
}

type EntrypointHandler func(input map[string]interface{}) (interface{}, error)

func buildHandler(entrypointHandler EntrypointHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Debug("Perform request received")

		var input map[string]interface{}
		err := json.NewDecoder(req.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := entrypointHandler(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(output)
	}
}
