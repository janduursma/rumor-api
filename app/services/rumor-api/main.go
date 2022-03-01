package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go.uber.org/zap/zapcore"

	"github.com/emadolsky/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

var build = "develop"

func main() {
	// Construct the application logger.
	log, err := initLogger("RUMOR-API")
	if err != nil {
		log.Fatalf("error constructing logger: %s", err)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// Check what maxprocs reports.
	opt := maxprocs.Logger(log.Infof)

	// Set the correct number of threads for the rumor-api
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(opt); err != nil {
		return fmt.Errorf("error setting maxprocs: %w", err)
	}

	log.Infow("starting rumor-api", "version", build, "CPU", runtime.GOMAXPROCS(0))
	defer log.Infow("shutdown completed")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Infow("shutting down rumor-api")

	return nil
}

func initLogger(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": service,
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}
