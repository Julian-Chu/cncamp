package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var requestLatencyHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace: "httpserver",
	Name:      "request_latency",
	Help:      "HTTP request latency distributions.",
	Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
})
