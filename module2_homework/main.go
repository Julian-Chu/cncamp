package main

import (
	"net"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
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

	r := httprouter.New()
	r.GET("/", indexHandler)
	r.GET("/healthz", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})

	port := ":8000"
	log.Println("listening to port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	for k, vals := range r.Header {
		for _, val := range vals {
			w.Header().Add(k, val)
		}
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.WithError(err).Error("can' get ip addr")
	} else {
		log.WithFields(logrus.Fields{
			"ip":     ip,
			"status": http.StatusOK,
		}).Info("user:")
	}
	w.Header().Add(VERSION, version)
	w.WriteHeader(http.StatusOK)
}
