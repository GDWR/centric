package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func LoggerMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)

			log.Debug().
				Str("endpoint", r.Pattern).
				Dur("duration", time.Since(start)).
				Send()
		})
	}
}
