package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof" // HL
)

func main() {
	http.HandleFunc("/test", heavyHandler)
	if err := http.ListenAndServe(":6061", nil); err != nil {
		log.Fatal(err)
	}
}

func heavyHandler(w http.ResponseWriter, req *http.Request) {
	for i := 0; i < 1000000; i++ {
		math.Pow(78, 23)
	}
	fmt.Fprint(w, "Some heavy work done!")
}
