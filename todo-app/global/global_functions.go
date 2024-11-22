package global

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func HttpError(w http.ResponseWriter, logMessage string, httpMessage string, code int, err error) {
	log.Errorf("%s: %v", logMessage, err)

	errorResponse := struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	}{
		Error:  httpMessage,
		Status: code,
	}
	w.WriteHeader(code)
	errEncode := json.NewEncoder(w).Encode(errorResponse)
	if errEncode != nil {
		log.Errorf("Failed to encode JSON error response: %v", errEncode)
		http.Error(w, "An internal server error occurred", http.StatusInternalServerError)
	}
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

func StrPtr(s string) *string {
	return &s
}

func TimePtr(t time.Time) *time.Time {
	return &t
}

func IntPtr(i int) *int {
	return &i
}
