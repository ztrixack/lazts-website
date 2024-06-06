package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Moduler interface {
	Serve() error
	Shutdown(ctx context.Context) error
	Get(path string, handler http.HandlerFunc)
	StaticFileServer(prefix string, handler http.HandlerFunc)
	Use(middlewares ...func(http.Handler) http.Handler)
}

type module struct {
	Address     string
	config      *config
	router      *mux.Router
	server      *http.Server
	middlewares []func(http.Handler) http.Handler
}

func New() *module {
	c := parseConfig()
	r := mux.NewRouter()
	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)

	return &module{
		config: c,
		router: r,
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		Address: addr,
	}
}

func (m *module) Serve() error {
	return m.server.ListenAndServe()
}

func (m *module) Shutdown(ctx context.Context) error {
	return m.server.Shutdown(ctx)
}

func (m *module) Get(path string, handler http.HandlerFunc) {
	chain := m.applyMiddlewares(handler).(http.HandlerFunc)
	m.router.HandleFunc(path, chain).Methods(http.MethodGet)
}

func (m *module) StaticFileServer(prefix string, handler http.HandlerFunc) {
	m.router.PathPrefix(prefix).Handler(http.StripPrefix(prefix, handler))
}

func (m *module) Use(middlewares ...func(http.Handler) http.Handler) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *module) applyMiddlewares(handler http.Handler) http.Handler {
	for _, middleware := range m.middlewares {
		handler = middleware(handler)
	}
	return handler
}
