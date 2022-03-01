package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/emadolsky/automaxprocs/maxprocs"
)

var build = "develop"

func main() {
	// Set the correct number of threads for the rumor-api
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		log.Fatalf("error setting maxprocs: %s", err)
	}

	log.Printf("starting rumor-api, build[%s] CPU[%d]", build, runtime.GOMAXPROCS(0))
	defer log.Println("rumor-api ended")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("stopping rumor-api")
}
