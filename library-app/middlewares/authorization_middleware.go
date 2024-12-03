package middlewares

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ErrorJSON struct {
	Message string `json:"db_error"`
}

func WriteErrorJSON(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)

	errJson := ErrorJSON{errorMessage}

	err := json.NewEncoder(w).Encode(errJson)
	if err != nil {
		log.Errorf("[WriteErrorJSON] Error encoding db_error response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer abc" {
			log.Errorf("[AuthorizationMiddleware] Unauthorized access attempt: missing or invalid Authorization header: %s", authHeader)

			WriteErrorJSON(w, http.StatusUnauthorized,
				fmt.Sprintf("Unauthorized access attempt: missing or invalid Authorization header: %s",
					authHeader))
			return
		}
		next.ServeHTTP(w, r)
	})
}
