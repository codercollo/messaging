package http

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applicatiion/json: charset=UTF-8")

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"methods": r.Method,
				"path":    r.URL.Path,
			}).Info("handled request")

		next.ServeHTTP(w, r)
	})
}
