package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metricRequestInternalErrors = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "request_internal_errors",
			Help: "Number of server errors when calling a Job"})
)

var (
	metricRequestsStarted = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_started",
			Help: "Total number of started requests calling Job (may not be finished yet)"})
)

var (
	metricRequestsDone = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_done",
			Help: "Total number of finished requests calling Job (processed and done)"})
)

var (
	metricEndpointRequestsStarted = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "endpoint_requests_started",
			Help: "Total number of started requests calling Job (may not be finished yet)"})
)

var (
	metricRequestDuration = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "request_duration",
			Help:    "Duration of model call",
			Buckets: []float64{0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.075, 0.1, 0.25, 0.5, 0.75, 1.0, 2.5, 5.0, 7.5, 10.0, 25.0, 50.0, 75.0, 100.0, 250.0, 500.0, 750.0, 1000.0}}) //+Inf is implicit
)

var (
	metricLastCallTimestamp = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "last_call_timestamp",
			Help: "Timestamp (in seconds) of the last request calling Job"})
)

func MetricHandler() gin.HandlerFunc {
	handler := promhttp.Handler()

	return func(context *gin.Context) {
		handler.ServeHTTP(context.Writer, context.Request)
	}
}
