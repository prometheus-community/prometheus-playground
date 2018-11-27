package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port = os.Getenv("PORT")

	processedOps = promauto.NewCounter(prometheus.CounterOpts{
		Name:      "processed_ops_total",
		Help:      "The total number of processed events",
		Namespace: "myservice",
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
	address := fmt.Sprintf(":%s", port)

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(address, nil))
}
