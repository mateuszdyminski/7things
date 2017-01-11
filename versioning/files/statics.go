package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var version = flag.String("version", "1.0", "Application version")

func main() {
	flag.Parse()

	router := mux.NewRouter()

	staticsRoot := fmt.Sprintf("/s/%s/", *version)
	staticsPath := staticsRoot + "{path:.*}"
	router.Handle(staticsPath, http.StripPrefix(staticsRoot, http.FileServer(http.Dir("statics"))))
	router.Handle("/s/{path:.*}", http.StripPrefix("/s/", http.FileServer(http.Dir("statics"))))

	http.ListenAndServe(":8082", router)
}
