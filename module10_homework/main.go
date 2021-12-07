package main

import (
	"context"
	"flag"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

const VERSION = "VERSION"

var log = logrus.New()
var version = os.Getenv(VERSION)

func main() {
	if version == "" {
		version = "0.0.0"
	}

	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  false,
		//DisableColors: true,
	},
	)

	prometheus.MustRegister(requestLatencyHistogram)

	r := httprouter.New()
	r.GET("/", indexHandler)
	r.GET("/healthz", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})

	port := ":8000"
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}
	log.Println("listening to port", port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("server err: %v", err)
		}
		log.Println("server is closed")
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("shut down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server err: %v", err)
	}
}

var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	delay := randGenerator.Intn(2000)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	for k, vals := range r.Header {
		for _, val := range vals {
			w.Header().Add(k, val)
		}
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.WithError(err).Error("can't get ip addr")
	} else {
		log.WithFields(logrus.Fields{
			"ip":     ip,
			"status": http.StatusOK,
		}).Info("user:")
	}
	w.Header().Add(VERSION, version)
	w.WriteHeader(http.StatusOK)
}

var (
	normDomain = flag.Float64("normal.domain", 0.0002, "The domain for the normal distribution.")
	normMean   = flag.Float64("normal.mean", 0.00001, "The mean for the normal distribution.")
)

var requestLatencyHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "httpserver",
	Help:    "HTTP request latency distributions.",
	Buckets: prometheus.LinearBuckets(*normMean-5**normDomain, .5**normDomain, 20),
})
