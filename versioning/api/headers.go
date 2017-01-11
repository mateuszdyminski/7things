package main

import (
	"net/http"
	"strings"
	"strconv"
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

var accept2ver = map[string]int{
	"vnd.testapp.v1": RestAPIv1,
	"vnd.testapp.v2": RestAPIv2,
}

// AcceptedVersion checks Accept header in request to parse API version.
func AcceptedVersion(req *http.Request) int {
	a := strings.Split(req.Header.Get("Accept"), "/")
	var v string
	if len(a) > 1 {
		versionAndType := strings.Split(a[1], "+")

		v = versionAndType[0]
		if len(versionAndType) > 1 {
			req.Header.Set("X-Version", strconv.Itoa(accept2ver[v]))
			req.Header.Set("X-Content", versionAndType[1])
		}
	}

	version := accept2ver[v]
	if version < 1 {
		return DefaultAPIVersion
	}

	return version
}

// Ver - middleware to get version of request.
func Ver(endpointHandler func(http.ResponseWriter, *http.Request, int)) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		endpointHandler(res, req, AcceptedVersion(req)) // HL
	})
}

func testHandler(resp http.ResponseWriter, req *http.Request, v int) {
	if v == RestAPIv2 {
		fmt.Fprint(resp, "Api V2")
	} else {
		fmt.Fprint(resp, "Api V1")
	}
}