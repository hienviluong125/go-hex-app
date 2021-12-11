package middleware

import (
	"encoding/json"
	"hienviluong125/go-hex-app/errorhandler"
	"hienviluong125/go-hex-app/logger"
	"net/http"
)

func HandleErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Add("Content-type", "application/json")

				if appErr, ok := err.(*errorhandler.AppError); ok {
					w.WriteHeader(appErr.StatusCode)
					json.NewEncoder(w).Encode(appErr)
					return
				}

				appErr := errorhandler.ErrInternal(err.(error))
				w.WriteHeader(appErr.StatusCode)
				json.NewEncoder(w).Encode(appErr)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func HandleLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.InfoHttpRequest(r)
		next.ServeHTTP(w, r)
	})
}
