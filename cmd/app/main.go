package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Setup zerolog to log in pretty console format
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// Log a message
	log.Info().Msg("Hello, World!")

	go func() {
		log.Info().Msg("Starting application")
	}()

	// Channel to listen for interrupt or terminate signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sig)

	// Block until a signal is received.
	<-sig

	log.Info().Msg("Shutting down application...")
}
