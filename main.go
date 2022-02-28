package main

import (
	"github.com/emadolsky/automaxprocs/maxprocs"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var build = "develop"

func main() {
	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		log.Fatalf("error setting maxprocs: %w", err)
	}

	log.Printf("starting service, build[%s] CPU[%d]", build, runtime.GOMAXPROCS(0))
	defer log.Println("service ended")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("stopping service")
}
