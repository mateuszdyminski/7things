package main

import (
	"flag"
	"fmt"
	"github.com/mihasya/go-metrics-librato"
	"github.com/rcrowley/go-metrics"
	"net/http"
	"time"
)

var token = flag.String("token", "", "Librato token")
var hostname = flag.String("hostname", "md", "Hostname")

func main() {
	flag.Parse()
	// GC stats
	metrics.RegisterDebugGCStats(metrics.DefaultRegistry) // HL
	go metrics.CaptureDebugGCStats(metrics.DefaultRegistry, time.Minute) // HL

	// Memory stats
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry) // HL
	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, time.Minute) // HL

	// Run sender
	go librato.Librato(metrics.DefaultRegistry, // HL
		time.Minute, "mateusz.dyminski@gmail.com", *token, *hostname,
		[]float64{0.5, 0.85, 0.95, 0.99}, time.Millisecond,
	)
	// create custom counter
	errors := metrics.NewCounter() // HL
	metrics.Register("app.http.error", errors) // HL

	// configure http server
	http.HandleFunc("/error", func(w http.ResponseWriter, req *http.Request) {
		errors.Inc(1) // HL
		fmt.Fprint(w, "Error +1!")
	})
	http.ListenAndServe(":7001", nil)
}
