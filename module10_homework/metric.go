package main

import (
	"flag"

	"github.com/prometheus/client_golang/prometheus"
)

var requestLatencyHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace: "httpserver",
	Name:      "request_latency",
	Help:      "HTTP request latency distributions.",
	Buckets:   prometheus.LinearBuckets(*normMean-5**normDomain, .5**normDomain, 20),
})

var (
	normDomain = flag.Float64("normal.domain", 0.0002, "The domain for the normal distribution.")
	normMean   = flag.Float64("normal.mean", 0.00001, "The mean for the normal distribution.")
)
