package validation

import (
	"awesomeProject/library-app/models"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestValidateListMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		list           models.List
		expectedStatus int
		expectedError  string
	}{
		{
			name: "Valid List",
			list: models.List{
				ID:           uuid.New(),
				Name:         "Valid List Name",
				Description:  "A valid list description",
				CreationTime: time.Now(),
				Items: []models.Item{
					{
						ID:           uuid.New(),
						ListID:       uuid.New(),
						Title:        "Valid Item Title",
						Description:  "A valid item description",
						Tags:         []models.Tag{{Name: "Tag1"}},
						Completed:    false,
						CreationTime: time.Now(),
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Empty List Name",
			list: models.List{
				ID:          uuid.New(),
				Name:        "",
				Description: "Description",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "list name cannot be empty",
		},
		{
			name: "List Name with Whitespace",
			list: models.List{
				ID:          uuid.New(),
				Name:        " Invalid List Name ",
				Description: "Description",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "list name cannot begin or end with whitespace",
		},
		{
			name: "List Name Exceeds Length Limit",
			list: models.List{
				ID:          uuid.New(),
				Name:        strings.Repeat("a", 101),
				Description: "Description",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "list name cannot exceed 100 characters",
		},
		{
			name: "List Description Exceeds Length Limit",
			list: models.List{
				ID:          uuid.New(),
				Name:        "Valid List Name",
				Description: strings.Repeat("a", 256),
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "list description cannot exceed 255 characters",
		},
		{
			name: "With Items",
			list: models.List{
				ID:   uuid.New(),
				Name: "Valid List Name",
				Items: []models.Item{
					{
						ID:          uuid.New(),
						ListID:      uuid.New(),
						Title:       "Title",
						Description: "Description",
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "No Items",
			list: models.List{
				ID:   uuid.New(),
				Name: "Valid List Name",
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.list)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/lists", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			handler := ValidateListMiddleware(nextHandler)
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedStatus != http.StatusOK {
				responseBody := rec.Body.String()
				assert.Contains(t, responseBody, tt.expectedError)
			}
		})
	}
}
