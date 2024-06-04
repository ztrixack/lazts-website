package middlewares

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Str("method", r.Method).Str("url", r.URL.Path).Msg("request received")
		next.ServeHTTP(w, r)
		log.Debug().Msg("request handled successfully")
	})
}
