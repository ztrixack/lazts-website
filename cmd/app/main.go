package main

import (
	"context"
	"lazts/internal/handlers/api"
	"lazts/internal/handlers/books"
	"lazts/internal/handlers/file"
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/modules/imaging"
	"lazts/internal/modules/log"
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
)

func main() {
	server := http.New()
	server.Use(middlewares.Logging)

	imaging := imaging.New()
	markdown := markdown.New()

	webber := web.New(markdown)
	booker := book.New()

	books.New(server, webber, booker)
	file.New(server, watermark.New(imaging))
	api.New(server, webber, booker, vacation.New(markdown), memo.New(markdown))

	go func() {
		log.I("starting server on port %s", server.Address)
		if err := server.Serve(); err != nil {
			log.Err(err).C("failed to start server")
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sig)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Err(err).E("failed to gracefully shutdown server")
		return
	}

	log.I("shutting down application...")
}
