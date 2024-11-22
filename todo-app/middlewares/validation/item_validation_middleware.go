package validation

import (
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/models"
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func ValidateItemMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("[ValidateItem] Validating Item")
		var item models.Item

		bodyBytes := global.ReadBody(w, r, "ValidateItemMiddleware")

		err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&item)
		if err != nil {
			global.HttpError(
				w,
				"[ValidateItemMiddleware] Failed to decode JSON",
				"Invalid JSON format in request body",
				http.StatusBadRequest,
				err,
			)
			return
		}

		err = validateItemFields(item)
		if err != nil {
			global.HttpError(
				w,
				"[ValidateItemMiddleware] validation failed",
				err.Error(),
				http.StatusBadRequest,
				err,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateItemFields(item models.Item) error {
	if item.Title == "" {
		return errors.New("item title cannot be empty")
	}

	if strings.Trim(item.Title, " \t\n") != item.Title {
		return errors.New("item title cannot begin or end with whitespace")
	}

	if len(item.Title) > 100 {
		return errors.New("item title cannot exceed 100 characters")
	}

	if len(item.Description) > 255 {
		return errors.New("item description cannot exceed 255 characters")
	}

	for _, tag := range item.Tags {
		err := validateTagFields(tag)
		if err != nil {
			return errors.New("invalid tag: " + err.Error())
		}
	}

	return nil
}
