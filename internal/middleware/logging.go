package middleware

import (
    "net/http"
    "time"
    "project/pkg/logger"
)

func LoggingMiddleware(logger *logger.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            next.ServeHTTP(w, r)
            logger.LogRequest(r.Method, r.URL.Path, time.Since(start))
        })
    }
}
