package errors

import (
	"encoding/json"
	"errors"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func HttpError(w http.ResponseWriter, logMessage string, httpMessage string, code int, err error) {
	log.Errorf("%s: %v", logMessage, err)

	errorResponse := ErrorResponse{
		Error:  httpMessage,
		Status: code,
	}
	w.WriteHeader(code)
	errEncode := json.NewEncoder(w).Encode(errorResponse)
	if errEncode != nil {
		log.Errorf("Failed to encode JSON db response: %v", errEncode)
		http.Error(w, "An internal server db occurred", http.StatusInternalServerError)
	}
}

func HttpDBError(w http.ResponseWriter, err error) {
	var (
		code        int
		httpMessage string
		logMessage  string
	)

	var dbError *db.DBError
	if !errors.As(err, &dbError) {
		dbError = db.NewDBError(db.InternalError, "Error is not db error")
	}

	switch dbError.Type {
	case db.ValidationError:
		code = http.StatusBadRequest
		httpMessage = dbError.Error()
		logMessage = "Validation Error"
	case db.NotFoundError:
		code = http.StatusNotFound
		httpMessage = "Resource not found"
		logMessage = "Not Found Error"
	case db.InternalError:
		code = http.StatusInternalServerError
		httpMessage = "An internal server error occurred"
		logMessage = "Internal Server Error"
	default:
		code = http.StatusInternalServerError
		httpMessage = "An unknown error occurred"
		logMessage = "Unknown Error"
	}

	log.Errorf("%s: %v", logMessage, dbError.Error())

	errorResponse := ErrorResponse{
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
