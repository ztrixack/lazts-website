package middlewares

import (
	"net/http"

	"lazts/internal/modules/log"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Fields("method", r.Method, "url", r.URL.Path, "query", r.URL.Query()).D("request received")
		next.ServeHTTP(w, r)
		log.Fields("method", r.Method, "url", r.URL.Path, "query", r.URL.Query()).D("request success")
	})
}
