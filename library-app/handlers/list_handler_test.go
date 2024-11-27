package handlers

import (
	"awesomeProject/library-app/managers/interfaces/automock"
	"awesomeProject/library-app/models"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupListHandlerTests() (*ListHandler, *automock.ListManager) {
	mockListManager := &automock.ListManager{}
	listHandler := NewListHandler(mockListManager)
	return listHandler, mockListManager
}

func TestListHandler_GetAll_Success(t *testing.T) {
	listHandler, mockListManager := setupListHandlerTests()

	lists := []models.List{
		{ID: uuid.New(), Name: "List1", Description: "First list"},
		{ID: uuid.New(), Name: "List2", Description: "Second list"},
	}
	mockListManager.On("GetAll").Return(lists)

	req := httptest.NewRequest(http.MethodGet, "/api/lists", nil)
	rec := httptest.NewRecorder()

	listHandler.GetAll(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result []models.List
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, lists, result)

	mockListManager.AssertExpectations(t)
}

func TestListHandler_Get_Success(t *testing.T) {
	listHandler, mockListManager := setupListHandlerTests()

	listID := uuid.New()
	list := models.List{ID: listID, Name: "List1", Description: "First list"}
	mockListManager.On("Get", listID).Return(list, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/lists/"+listID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": listID.String()})
	rec := httptest.NewRecorder()

	listHandler.Get(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.List
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, list, result)

	mockListManager.AssertExpectations(t)
}

func TestListHandler_Get_NotFound(t *testing.T) {
	listHandler, mockListManager := setupListHandlerTests()

	nonexistentID := uuid.New()
	mockListManager.On("Get", nonexistentID).Return(models.List{}, errors.New("not found"))

	req := httptest.NewRequest(http.MethodGet, "/api/lists/"+nonexistentID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": nonexistentID.String()})
	rec := httptest.NewRecorder()

	listHandler.Get(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	mockListManager.AssertExpectations(t)
}

func TestListHandler_Create_Success(t *testing.T) {
	listHandler, mockListManager := setupListHandlerTests()

	newList := models.List{Name: "New List", Description: "A new list"}
	createdList := newList
	createdList.ID = uuid.New()
	mockListManager.On("Create", newList).Return(createdList, nil)

	body, _ := json.Marshal(newList)
	req := httptest.NewRequest(http.MethodPost, "/api/lists", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	listHandler.Create(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.List
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, createdList, result)

	mockListManager.AssertExpectations(t)
}

func TestListHandler_Create_InvalidJSON(t *testing.T) {
	listHandler, mockListManager := setupListHandlerTests()

	req := httptest.NewRequest(http.MethodPost, "/api/lists", bytes.NewBuffer([]byte("invalid json")))
	rec := httptest.NewRecorder()

	listHandler.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	mockListManager.AssertExpectations(t)
}

func TestListHandler_Update_Success(t *testing.T) {
	listHandler, mockListManager := setupListHandlerTests()

	listID := uuid.New()
	updateList := models.List{ID: listID, Name: "UpdatedList", Description: "Updated description"}
	mockListManager.On("Update", updateList).Return(updateList, nil)

	body, _ := json.Marshal(updateList)
	req := httptest.NewRequest(http.MethodPut, "/api/lists/"+listID.String(), bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": listID.String()})
	rec := httptest.NewRecorder()

	listHandler.Update(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.List
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, updateList, result)

	mockListManager.AssertExpectations(t)
}

func TestListHandler_Update_InvalidUUID(t *testing.T) {
	listHandler, _ := setupListHandlerTests()

	req := httptest.NewRequest(http.MethodPut, "/api/lists/invalid-uuid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid-uuid"})
	rec := httptest.NewRecorder()

	listHandler.Update(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestListHandler_Delete_Success(t *testing.T) {
	listHandler, mockListManager := setupListHandlerTests()

	listID := uuid.New()
	deletedList := models.List{ID: listID, Name: "ListToDelete", Description: "A list to be deleted"}
	mockListManager.On("Delete", listID).Return(deletedList, nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/lists/"+listID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": listID.String()})
	rec := httptest.NewRecorder()

	listHandler.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.List
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, deletedList, result)

	mockListManager.AssertExpectations(t)
}

func TestListHandler_Delete_NotFound(t *testing.T) {
	listHandler, mockListManager := setupListHandlerTests()

	nonexistentID := uuid.New()
	mockListManager.On("Delete", nonexistentID).Return(models.List{}, errors.New("not found"))

	req := httptest.NewRequest(http.MethodDelete, "/api/lists/"+nonexistentID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": nonexistentID.String()})
	rec := httptest.NewRecorder()

	listHandler.Delete(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	mockListManager.AssertExpectations(t)
}