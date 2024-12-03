package validation

import (
	"awesomeProject/library-app/global"
	"awesomeProject/library-app/models"
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func ValidateTagMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("[ValidateTagMiddleware] Validating Tag")
		var tag models.Tag

		bodyBytes := global.ReadBody(w, r, "ValidateTagMiddleware")

		err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&tag)
		if err != nil {
			global.HttpError(
				w,
				"[ValidateTagMiddleware] Failed to decode JSON",
				"Invalid JSON format in request body",
				http.StatusBadRequest,
				err,
			)
			return
		}

		err = validateTagFields(tag)
		if err != nil {
			global.HttpError(
				w,
				"[ValidateTagMiddleware] validation failed",
				err.Error(),
				http.StatusBadRequest,
				err,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateTagFields(tag models.Tag) error {
	if tag.Name == "" {
		return errors.New("tag name cannot be empty")
	}

	if strings.ContainsAny(tag.Name, " \t\n") {
		return errors.New("tag name cannot contain whitespace")
	}

	if len(tag.Name) > 100 {
		return errors.New("tag name cannot exceed 100 characters")
	}

	return nil
}
