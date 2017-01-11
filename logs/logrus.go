package main

import (
	log "github.com/Sirupsen/logrus" // HL
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{ // HL
		"lang": "golang",
	}).Info("Hello Wroclaw")
}
