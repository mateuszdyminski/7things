package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

// Variables injected by LDFLAGS -X flag
var appVersion string = "unknown"
var lastCommitTime string = "unknown"
var lastCommitHash string = "unknown"
var lastCommitUser string = "unknown"
var buildTime string = "unknown"

// Globals used in healthz
var hostname string = "unknown"
var startedAt time.Time = time.Now().UTC()
var build Build

// HealthzResponse holds extended health check status.
type HealthzResponse struct {
	Build     Build             `json:"build"`
	Hostname  string            `json:"hostname"`
	Uptime    string            `json:"uptime"`
	StartedAt string            `json:"startedAt"`
	DbStatus  map[string]string `json:"dbStatuse"`
}

// Build holds information about the build.
type Build struct {
	Version    string `json:"version"`
	BuildTime  string `json:"buildTime"`
	LastCommit Commit `json:"lastCommit"`
}

// Commit holds information about last git commit.
type Commit struct {
	Id     string `json:"id"`
	Time   string `json:"time"`
	Author string `json:"author"`
}

func main() {
	build = Build{
		Version:   appVersion,
		BuildTime: buildTime,
		LastCommit: Commit{
			Author: lastCommitUser,
			Id:     lastCommitHash,
			Time:   lastCommitTime,
		},
	}

	hostname, _ = os.Hostname() // HL

	http.HandleFunc("/health", healthz) // HL
	if err := http.ListenAndServe(":9005", nil); err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	resp := HealthzResponse{
		Hostname:  hostname,
		StartedAt: startedAt.Format("2006-01-02_15:04:05"),
		Uptime:    time.Now().UTC().Sub(startedAt).String(),
		Build:     build,
		DbStatus:  checkDbStatus(), // HL
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't marshal json", http.StatusBadRequest)
		return
	}

	w.Write(json)
}

func checkDbStatus() map[string]string {
	results := make(map[string]string)

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello") // HL
	if err != nil {
		results["status"] = fmt.Sprintf("failed with error: %v", err)
		return results
	}
	defer db.Close()

	err = db.Ping() // HL
	if err != nil {
		results["status"] = fmt.Sprintf("failed with error: %v", err)
		return results
	}

	results["status"] = "OK"

	return results
}
