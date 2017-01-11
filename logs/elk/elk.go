package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/bshuster-repo/logrus-logstash-hook"
)

func main() {
	log := logrus.New()
	hook, err := logrus_logstash.NewHook("tcp", "localhost:5000", "test") // HL

	if err != nil {
		log.Fatal(err)
	}
	log.Hooks.Add(hook) // HL

	ctx := log.WithFields(logrus.Fields{
		"lang": "golang",
	})

	ctx.Info("Hello Wroclaw!")
}
