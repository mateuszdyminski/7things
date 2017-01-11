package main

import "log"

// Variables injected by -X flag
var appVersion string = "unknown"
var lastCommitTime string = "unknown"
var lastCommitHash string = "unknown"
var lastCommitUser string = "unknown"
var buildTime string = "unknown"

func main() {
	log.Printf("Version: %s \n", appVersion)
	log.Printf("Time of last commit: %s \n", lastCommitTime)
	log.Printf("Hash of last commit: %s \n", lastCommitHash)
	log.Printf("User of last commit: %s \n", lastCommitUser)
	log.Printf("Build time: %s \n", buildTime)
}
