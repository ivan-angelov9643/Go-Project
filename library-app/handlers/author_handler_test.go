package handlers

import (
	"awesomeProject/library-app/managers/interfaces/automock"
	"awesomeProject/library-app/models"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupAuthorHandlerTests() (*AuthorHandler, *automock.AuthorManager) {
	mockAuthorManager := &automock.AuthorManager{}
	authorHandler := NewAuthorHandler(mockAuthorManager)
	return authorHandler, mockAuthorManager
}

func TestAuthorHandler_GetAll_Success(t *testing.T) {
	authorHandler, mockAuthorManager := setupAuthorHandlerTests()

	authors := []models.Author{
		{
			BaseModel:   models.BaseModel{ID: uuid.New()},
			FirstName:   "John",
			LastName:    "Doe",
			Nationality: "American",
		},
		{
			BaseModel:   models.BaseModel{ID: uuid.New()},
			FirstName:   "Jane",
			LastName:    "Smith",
			Nationality: "British",
		},
	}
	mockAuthorManager.On("GetAll").Return(authors, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/authors", nil)
	rec := httptest.NewRecorder()

	authorHandler.GetAll(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result []models.Author
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, authors, result)

	mockAuthorManager.AssertExpectations(t)
}

func TestAuthorHandler_Get_Success(t *testing.T) {
	authorHandler, mockAuthorManager := setupAuthorHandlerTests()

	authorID := uuid.New()
	author := models.Author{
		BaseModel:   models.BaseModel{ID: authorID},
		FirstName:   "John",
		LastName:    "Doe",
		Nationality: "American",
	}
	mockAuthorManager.On("Get", authorID).Return(author, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/authors/"+authorID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": authorID.String()})
	rec := httptest.NewRecorder()

	authorHandler.Get(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.Author
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, author, result)

	mockAuthorManager.AssertExpectations(t)
}

func TestAuthorHandler_Get_InvalidUUID(t *testing.T) {
	authorHandler, _ := setupAuthorHandlerTests()

	req := httptest.NewRequest(http.MethodGet, "/api/authors/invalid-uuid", nil)
	rec := httptest.NewRecorder()

	authorHandler.Get(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestAuthorHandler_Create_Success(t *testing.T) {
	authorHandler, mockAuthorManager := setupAuthorHandlerTests()

	newAuthor := models.Author{
		FirstName:   "John",
		LastName:    "Doe",
		Nationality: "American",
	}
	createdAuthor := newAuthor
	createdAuthor.ID = uuid.New()

	mockAuthorManager.On("Create", newAuthor).Return(createdAuthor, nil)

	body, _ := json.Marshal(newAuthor)
	req := httptest.NewRequest(http.MethodPost, "/api/authors", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	authorHandler.Create(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.Author
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, createdAuthor, result)

	mockAuthorManager.AssertExpectations(t)
}

func TestAuthorHandler_Create_InvalidJSON(t *testing.T) {
	authorHandler, _ := setupAuthorHandlerTests()

	req := httptest.NewRequest(http.MethodPost, "/api/authors", bytes.NewBuffer([]byte("invalid json")))
	rec := httptest.NewRecorder()

	authorHandler.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestAuthorHandler_Update_Success(t *testing.T) {
	authorHandler, mockAuthorManager := setupAuthorHandlerTests()

	authorID := uuid.New()
	updatedAuthor := models.Author{
		BaseModel:   models.BaseModel{ID: authorID},
		FirstName:   "John",
		LastName:    "Smith",
		Nationality: "American",
	}
	mockAuthorManager.On("Update", updatedAuthor).Return(updatedAuthor, nil)

	body, _ := json.Marshal(updatedAuthor)
	req := httptest.NewRequest(http.MethodPut, "/api/authors/"+authorID.String(), bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": authorID.String()})
	rec := httptest.NewRecorder()

	authorHandler.Update(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.Author
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, updatedAuthor, result)

	mockAuthorManager.AssertExpectations(t)
}

func TestAuthorHandler_Delete_Success(t *testing.T) {
	authorHandler, mockAuthorManager := setupAuthorHandlerTests()

	authorID := uuid.New()
	deletedAuthor := models.Author{
		BaseModel:   models.BaseModel{ID: authorID},
		FirstName:   "John",
		LastName:    "Doe",
		Nationality: "American",
	}
	mockAuthorManager.On("Delete", authorID).Return(deletedAuthor, nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/authors/"+authorID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": authorID.String()})
	rec := httptest.NewRecorder()

	authorHandler.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.Author
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, deletedAuthor, result)

	mockAuthorManager.AssertExpectations(t)
}
