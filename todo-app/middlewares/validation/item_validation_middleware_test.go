package validation

import (
	"awesomeProject/todo-app/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateItemMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		item           models.Item
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Valid Item",
			item: models.Item{
				ID:           uuid.New(),
				ListID:       uuid.New(),
				Title:        "Valid Title",
				Description:  "A valid description",
				Tags:         []models.Tag{{Name: "Tag1"}},
				Completed:    false,
				CreationTime: time.Now(),
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Empty Title",
			item: models.Item{
				ID:          uuid.New(),
				ListID:      uuid.New(),
				Title:       "",
				Description: "Description",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item title cannot be empty",
		},
		{
			name: "Title with Whitespace",
			item: models.Item{
				ID:          uuid.New(),
				ListID:      uuid.New(),
				Title:       " Invalid Title ",
				Description: "Description",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item title cannot begin or end with whitespace",
		},
		{
			name: "Title Exceeds Length Limit",
			item: models.Item{
				ID:          uuid.New(),
				ListID:      uuid.New(),
				Title:       strings.Repeat("a", 101),
				Description: "Description",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item title cannot exceed 100 characters",
		},
		{
			name: "Description Exceeds Length Limit",
			item: models.Item{
				ID:          uuid.New(),
				ListID:      uuid.New(),
				Title:       "Valid Title",
				Description: strings.Repeat("a", 256),
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "item description cannot exceed 255 characters",
		},
		{
			name: "With Tags",
			item: models.Item{
				ID:          uuid.New(),
				ListID:      uuid.New(),
				Title:       "Valid Title",
				Description: "Description",
				Tags:        []models.Tag{{Name: "Tag1"}},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "No Tags",
			item: models.Item{
				ID:          uuid.New(),
				ListID:      uuid.New(),
				Title:       "Valid Title",
				Description: "Description",
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.item)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/items", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			handler := ValidateItemMiddleware(nextHandler)
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedStatus != http.StatusOK {
				responseBody := rec.Body.String()
				assert.Contains(t, responseBody, tt.expectedError)
			}
		})
	}
}
