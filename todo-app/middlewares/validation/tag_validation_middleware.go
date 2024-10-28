package validation

import (
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/structs"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func ValidateTag(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("[ValidateTag] Validating Tag")
		var tag structs.Tag

		err := json.NewDecoder(r.Body).Decode(&tag)
		if err != nil {
			global.HttpError(
				w,
				"[ValidateTag] Failed to decode JSON",
				"Invalid JSON format in request body",
				http.StatusBadRequest,
				err,
			)
			return
		}

		if err := validateTagName(tag.Name); err != nil {
			global.HttpError(
				w,
				"[ValidateTag] Name validation failed",
				err.Error(),
				http.StatusBadRequest,
				err,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validateTagName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("tag name cannot be empty")
	}
	if len(name) > 100 {
		return fmt.Errorf("tag name cannot exceed 100 characters")
	}
	return nil
}
