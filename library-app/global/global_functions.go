package global

import (
	"awesomeProject/library-app/global/db"
	"encoding/json"
	"errors"
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

//func ReadBody(w http.ResponseWriter, r *http.Request, functionCalled string) []byte {
//	bodyBytes, err := io.ReadAll(r.Body)
//	if err != nil {
//		HttpError(
//			w,
//			"["+functionCalled+"] Failed to read request body",
//			err.Error(),
//			http.StatusInternalServerError,
//			err,
//		)
//		return nil
//	}
//
//	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
//	return bodyBytes
//}

func StrPtr(s string) *string {
	return &s
}

//func TimePtr(t time.Time) *time.Time {
//	return &t
//}

func IntPtr(i int) *int {
	return &i
}

//func FloatPtr(i float64) *float64 {
//	return &i
//}

func IsGlobal(request *http.Request) bool {
	return request.Context().Value(GLOBAL_SCOPE) != nil
}

func GetOwnerID(request *http.Request) string {
	return request.Context().Value(CURRENT_USER_ID).(string)
}
