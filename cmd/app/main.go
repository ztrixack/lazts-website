package main

import (
	"context"
	"lazts/internal/handlers/api"
	"lazts/internal/handlers/file"
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/modules/imaging"
	"lazts/internal/modules/markdown"
	"lazts/internal/services/book"
	"lazts/internal/services/memo"
	"lazts/internal/services/vacation"
	"lazts/internal/services/watermark"
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

	server := http.New()
	server.Use(middlewares.Logging)

	imaging := imaging.New()
	markdown := markdown.New()

	file.New(server, watermark.New(imaging))
	api.New(server, web.New(markdown), book.New(), vacation.New(markdown), memo.New(markdown))

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
