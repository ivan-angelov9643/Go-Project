package handlers

import (
	"awesomeProject/todo-app/managers/interfaces/automock"
	"awesomeProject/todo-app/models"
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

func setupItemHandlerTests() (*ItemHandler, *automock.ItemManager) {
	mockItemManager := &automock.ItemManager{}
	itemHandler := NewItemHandler(mockItemManager)
	return itemHandler, mockItemManager
}

func TestItemHandler_GetAll_Success(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	items := []models.Item{
		{ID: uuid.New(), Title: "Item1", Completed: false},
		{ID: uuid.New(), Title: "Item2", Completed: true},
	}
	mockItemManager.On("GetAll").Return(items)

	req := httptest.NewRequest(http.MethodGet, "/api/items", nil)
	rec := httptest.NewRecorder()

	itemHandler.GetAll(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result []models.Item
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, items, result)

	mockItemManager.AssertExpectations(t)
}

func TestItemHandler_Get_Success(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	itemID := uuid.New()
	item := models.Item{ID: itemID, Title: "Item1", Completed: false}
	mockItemManager.On("Get", itemID).Return(item, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/items/"+itemID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": itemID.String()})
	rec := httptest.NewRecorder()

	itemHandler.Get(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.Item
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, item, result)

	mockItemManager.AssertExpectations(t)
}

func TestItemHandler_Get_NotFound(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	nonexistentID := uuid.New()
	mockItemManager.On("Get", nonexistentID).Return(models.Item{}, errors.New("not found"))

	req := httptest.NewRequest(http.MethodGet, "/api/items/"+nonexistentID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": nonexistentID.String()})
	rec := httptest.NewRecorder()

	itemHandler.Get(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	mockItemManager.AssertExpectations(t)
}

func TestItemHandler_Create_Success(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	newItem := models.Item{Title: "NewItem", Completed: false}
	createdItem := newItem
	createdItem.ID = uuid.New()
	mockItemManager.On("Create", newItem).Return(createdItem, nil)

	body, _ := json.Marshal(newItem)
	req := httptest.NewRequest(http.MethodPost, "/api/items", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	itemHandler.Create(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.Item
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, createdItem, result)

	mockItemManager.AssertExpectations(t)
}

func TestItemHandler_Create_InvalidJSON(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	req := httptest.NewRequest(http.MethodPost, "/api/items", bytes.NewBuffer([]byte("invalid json")))
	rec := httptest.NewRecorder()

	itemHandler.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	mockItemManager.AssertExpectations(t)
}

func TestItemHandler_Update_Success(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	itemID := uuid.New()
	updateItem := models.Item{ID: itemID, Title: "UpdatedItem", Completed: true}
	mockItemManager.On("Update", updateItem).Return(updateItem, nil)

	body, _ := json.Marshal(updateItem)
	req := httptest.NewRequest(http.MethodPut, "/api/items/"+itemID.String(), bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": itemID.String()})
	rec := httptest.NewRecorder()

	itemHandler.Update(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.Item
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, updateItem, result)

	mockItemManager.AssertExpectations(t)
}

func TestItemHandler_Update_InvalidUUID(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	req := httptest.NewRequest(http.MethodPut, "/api/items/invalid-uuid", nil)
	rec := httptest.NewRecorder()

	itemHandler.Update(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	mockItemManager.AssertExpectations(t)
}

func TestItemHandler_Delete_Success(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	itemID := uuid.New()
	deletedItem := models.Item{ID: itemID, Title: "DeletedItem", Completed: true}
	mockItemManager.On("Delete", itemID).Return(deletedItem, nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/items/"+itemID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": itemID.String()})
	rec := httptest.NewRecorder()

	itemHandler.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result models.Item
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, deletedItem, result)

	mockItemManager.AssertExpectations(t)
}

func TestItemHandler_Delete_NotFound(t *testing.T) {
	itemHandler, mockItemManager := setupItemHandlerTests()

	nonexistentID := uuid.New()
	mockItemManager.On("Delete", nonexistentID).Return(models.Item{}, errors.New("not found"))

	req := httptest.NewRequest(http.MethodDelete, "/api/items/"+nonexistentID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": nonexistentID.String()})
	rec := httptest.NewRecorder()

	itemHandler.Delete(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	mockItemManager.AssertExpectations(t)
}
