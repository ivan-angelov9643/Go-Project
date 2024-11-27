package validation

import (
	"awesomeProject/library-app/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateTagMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		tag            models.Tag
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Valid Tag",
			tag:            models.Tag{Name: "ValidTagName"},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Empty Tag Name",
			tag:            models.Tag{Name: ""},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "tag name cannot be empty",
		},
		{
			name:           "Tag Name with Whitespace",
			tag:            models.Tag{Name: "Invalid Tag"},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "tag name cannot contain whitespace",
		},
		{
			name:           "Tag Name Exceeds 100 Characters",
			tag:            models.Tag{Name: strings.Repeat("a", 101)},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "tag name cannot exceed 100 characters",
		},
		{
			name:           "Invalid JSON Format",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid JSON format in request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error

			if tt.name == "Invalid JSON Format" {
				body = []byte("{invalid_json")
			} else {
				body, err = json.Marshal(tt.tag)
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/tags", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			handler := ValidateTagMiddleware(nextHandler)
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedStatus != http.StatusOK {
				responseBody := rec.Body.String()
				assert.Contains(t, responseBody, tt.expectedError)
			}
		})
	}
}
