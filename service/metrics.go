package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	activeConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Number of active connections",
		},
	)

	requestSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_size_bytes",
			Help: "HTTP request size in bytes",
		},
		[]string{"method", "endpoint"},
	)

	responseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_response_size_bytes",
			Help: "HTTP response size in bytes",
		},
		[]string{"method", "endpoint", "status"},
	)
)

func initMetrics() {
	prometheus.MustRegister(activeConnections)
	prometheus.MustRegister(requestSize)
	prometheus.MustRegister(responseSize)
}

