package global

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func HttpError(w http.ResponseWriter, logMessage string, httpMessage string, code int, err error) {
	log.Errorf("%s: %v", logMessage, err)
	http.Error(w, httpMessage, code)
}

func ReadBody(w http.ResponseWriter, r *http.Request, functionCalled string) []byte {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		HttpError(
			w,
			"["+functionCalled+"] Failed to read request body",
			err.Error(),
			http.StatusInternalServerError,
			err,
		)
		return nil
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}
