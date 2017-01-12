package main

import (
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"fmt"
)

const (
	// RestAPIv1 first version of Test api.
	RestAPIv1 = iota + 1
	RestAPIv2

	// DefaultAPIVersion default version set when no http Accept header
	DefaultAPIVersion = 1
)

func main() {
	router := mux.NewRouter()

	router.Handle("/api/test", Ver(testHandler)).Methods("GET") // HL

	http.ListenAndServe(":8081", router)
}

// START
var accept2ver = map[string]int{
	"vnd.testapp.v1": RestAPIv1,
	"vnd.testapp.v2": RestAPIv2,
}

func AcceptedVersion(req *http.Request) int {
	a := strings.Split(req.Header.Get("Accept"), "/") // HL
	var v string
	if len(a) > 1 {
		versionAndType := strings.Split(a[1], "+") // HL
		v = versionAndType[0] // HL
	}

	version := accept2ver[v] // HL
	if version < 1 {
		return DefaultAPIVersion
	}

	return version
}
// STOP

func Ver(endpointHandler func(http.ResponseWriter, *http.Request, int)) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		endpointHandler(res, req, AcceptedVersion(req)) // HL
	})
}

func testHandler(resp http.ResponseWriter, req *http.Request, v int) {
	if v == RestAPIv2 { // HL
		fmt.Fprint(resp, "Api V2")
	} else {
		fmt.Fprint(resp, "Api V1")
	}
}