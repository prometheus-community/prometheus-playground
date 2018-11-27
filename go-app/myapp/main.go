package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	port = 2112
)

var (
	processedOps = promauto.NewCounter(prometheus.CounterOpts{
		Name:      "processed_ops_total",
		Help:      "The total number of processed events",
		Namespace: "myapp",
	})
)

func recordMetrics() {
	go func() {
		for {
			processedOps.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	address := fmt.Sprintf(":%d", port)

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(address, nil))
}
