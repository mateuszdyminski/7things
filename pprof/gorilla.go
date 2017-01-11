package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"net/http/pprof" // HL
)

func main() {
	r := mux.NewRouter()

	attachProfiler(r) // HL

	r.HandleFunc("/test", testHandler)
	if err := http.ListenAndServe(":6062", r); err != nil {
		log.Fatal(err)
	}
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	for i := 0; i < 1000000; i++ {
		math.Pow(78, 23)
	}
	fmt.Fprint(w, "Some heavy work done!")
}

func attachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}
