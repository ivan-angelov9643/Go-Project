package validation

import (
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/structs"
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func ValidateListMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("[ValidateListMiddleware] Validating List")
		var list structs.List

		bodyBytes := global.ReadBody(w, r, "ValidateListMiddleware")

		err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&list)
		if err != nil {
			global.HttpError(
				w,
				"[ValidateListMiddleware] Failed to decode JSON",
				"Invalid JSON format in request body",
				http.StatusBadRequest,
				err,
			)
			return
		}

		err = validateListFields(list)
		if err != nil {
			global.HttpError(
				w,
				"[ValidateListMiddleware] validation failed",
				err.Error(),
				http.StatusBadRequest,
				err,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateListFields(list structs.List) error {
	if list.Name == "" {
		return errors.New("list name cannot be empty")
	}

	if strings.Trim(list.Name, " \t\n") != list.Name {
		return errors.New("list name cannot begin or end with whitespace")
	}

	if len(list.Name) > 100 {
		return errors.New("list name cannot exceed 100 characters")
	}

	if len(list.Description) > 255 {
		return errors.New("list description cannot exceed 255 characters")
	}

	for _, item := range list.Items {
		err := validateItemFields(item)
		if err != nil {
			return errors.New("invalid item: " + err.Error())
		}
	}

	return nil
}
