package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type Moduler interface {
	Serve() error
	Shutdown(c context.Context) error
	Get(path string, handler http.HandlerFunc)
	Use(middlewares ...func(http.Handler) http.Handler)
}

type module struct {
	Address     string
	config      *config
	mux         *http.ServeMux
	server      *http.Server
	middlewares []func(http.Handler) http.Handler
}

func New() *module {
	c := parseConfig()
	m := http.NewServeMux()
	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)
	return &module{
		config: c,
		mux:    m,
		server: &http.Server{
			Addr:    addr,
			Handler: m,
		},
		Address: addr,
	}
}

func (m *module) Serve() error {
	return m.server.ListenAndServe()
}

func (m *module) Shutdown(c context.Context) error {
	return m.server.Shutdown(c)
}

func (m *module) Get(path string, handler http.HandlerFunc) {
	m.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if !strings.HasPrefix(r.URL.Path, "/static") && r.URL.Path != path {
			log.Debug().Str("url", r.URL.Path).Str("path", path).Msg("invalid path")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var current http.Handler = handler
		for _, middleware := range m.middlewares {
			current = middleware(current)
		}

		current.ServeHTTP(w, r)
	})
}

func (m *module) Use(middlewares ...func(http.Handler) http.Handler) {
	m.middlewares = append(m.middlewares, middlewares...)
}
