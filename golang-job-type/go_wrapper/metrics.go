package main

import (
        "time"

        "github.com/gin-gonic/gin"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
        go func() {
                for {
                        opsProcessed.Inc()
                        time.Sleep(2 * time.Second)
                }
        }()
}

var (
        opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
                Name: "myapp_processed_ops_total",
                Help: "The total number of processed events",
        })
)

func MetricHandler() gin.HandlerFunc {
    recordMetrics()

    handler := promhttp.Handler()

    return func(context *gin.Context) {
        handler.ServeHTTP(context.Writer, context.Request)
    }
}
