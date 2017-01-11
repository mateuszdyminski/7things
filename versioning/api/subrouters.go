package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	apiV1 := api.PathPrefix("/v1").Subrouter()
	apiV2 := api.PathPrefix("/v2").Subrouter()

	test1 := apiV1.Path("/test").Subrouter()
	test1.Methods("GET").HandlerFunc(testV1Handler)

	test2 := apiV2.Path("/test").Subrouter()
	test2.Methods("GET").HandlerFunc(testV2Handler)

	http.ListenAndServe(":8080", r)
}

func testV1Handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "Api V1")
}

func testV2Handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "Api V2")
}
