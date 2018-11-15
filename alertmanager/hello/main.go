package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/unrolled/render"
)

const (
	port = 2112
)

var (
	renderer = render.New()
	address  = fmt.Sprintf(":%d", port)
)

func hello(w http.ResponseWriter, r *http.Request) {
	renderer.JSON(w, http.StatusOK, map[string]string{"hello": "world"})
}

func main() {
	router := chi.NewRouter()

	router.Get("/hello", hello)

	router.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(address, router))
}
