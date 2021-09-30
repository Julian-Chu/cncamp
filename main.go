package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

const VERSION = "VERSION"

func main() {
	r := httprouter.New()
	r.GET("/", indexHandler)
	r.GET("/healthz", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	for k, vals := range r.Header {
		for _, val := range vals {
			w.Header().Add(k, val)
		}
	}
	ver := os.Getenv(VERSION)
	if ver == "" {
		ver = "0.0.0"
	}
	w.Header().Add(VERSION, ver)
}
