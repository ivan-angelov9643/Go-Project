package middlewares

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SetJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("[SetJSONMiddleware] Request: Method=%s, Endpoint=%s. Setting Content-Type to application/json", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
