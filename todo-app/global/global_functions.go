package global

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HttpError(w http.ResponseWriter, logMessage string, httpMessage string, code int, err error) {
	log.Errorf("%s: %v", logMessage, err)
	http.Error(w, httpMessage, code)
}
