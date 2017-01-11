package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/square", squareVal)
	http.ListenAndServe(":9000", nil)
}

func squareVal(resp http.ResponseWriter, req *http.Request) {
	arg := req.URL.Query().Get("arg")
	if arg == "" {
		http.Error(resp, "arg should be set!", http.StatusBadRequest)
		return
	}

	val, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		http.Error(resp, fmt.Sprintf("can't parse arg! err: %v", err), http.StatusBadRequest)
		return
	}
	log.Printf("got arg: %f", val)

	result := math.Pow(val, 2)
	log.Printf("square value: %f", result)

	fmt.Fprint(resp, result)
}
