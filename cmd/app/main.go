package main

import (
	"context"
	"lazts/internal/handlers/api"
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/modules/markdown"
	"lazts/internal/services/web"
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

	markdown := markdown.New()
	server := http.New()
	server.Use(middlewares.Logging)

	web := web.New(markdown)

	api.New(server, web)

	go func() {
		log.Info().Msgf("starting server on port %s", server.Address)
		if err := server.Serve(); err != nil {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sig)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Err(err).Msg("failed to gracefully shutdown server")
		return
	}

	log.Info().Msg("shutting down application...")
}
