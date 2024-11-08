package handlers

import (
	"awesomeProject/todo-app/managers/interfaces/automock"
	"awesomeProject/todo-app/structs"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTagHandlerTests() (*TagHandler, *automock.TagManager) {
	mockTagManager := &automock.TagManager{}
	tagHandler := NewTagHandler(mockTagManager)
	return tagHandler, mockTagManager
}

func TestTagHandler_GetAll_Success(t *testing.T) {
	tagHandler, mockTagManager := setupTagHandlerTests()

	tags := []structs.Tag{
		{Name: "Tag1"},
		{Name: "Tag2"},
	}
	mockTagManager.On("GetAll").Return(tags)

	req := httptest.NewRequest(http.MethodGet, "/api/tags", nil)
	rec := httptest.NewRecorder()

	tagHandler.GetAll(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result []structs.Tag
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, tags, result)

	mockTagManager.AssertExpectations(t)
}

func TestTagHandler_Get_Success(t *testing.T) {
	tagHandler, mockTagManager := setupTagHandlerTests()

	tag := structs.Tag{Name: "Tag1"}
	mockTagManager.On("Get", "Tag1").Return(tag, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/tags/Tag1", nil)
	req = mux.SetURLVars(req, map[string]string{"name": "Tag1"})
	rec := httptest.NewRecorder()

	tagHandler.Get(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result structs.Tag
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, tag, result)

	mockTagManager.AssertExpectations(t)
}

func TestTagHandler_Get_NotFound(t *testing.T) {
	tagHandler, mockTagManager := setupTagHandlerTests()

	mockTagManager.On("Get", "UnknownTag").Return(structs.Tag{}, errors.New("not found"))

	req := httptest.NewRequest(http.MethodGet, "/api/tags/UnknownTag", nil)
	req = mux.SetURLVars(req, map[string]string{"name": "UnknownTag"})
	rec := httptest.NewRecorder()

	tagHandler.Get(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	mockTagManager.AssertExpectations(t)
}

func TestTagHandler_Create_Success(t *testing.T) {
	tagHandler, mockTagManager := setupTagHandlerTests()

	newTag := structs.Tag{Name: "NewTag"}
	mockTagManager.On("Create", newTag).Return(newTag, nil)

	body, _ := json.Marshal(newTag)
	req := httptest.NewRequest(http.MethodPost, "/api/tags", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	tagHandler.Create(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result structs.Tag
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, newTag, result)

	mockTagManager.AssertExpectations(t)
}

func TestTagHandler_Create_InvalidJSON(t *testing.T) {
	tagHandler, mockTagManager := setupTagHandlerTests()

	req := httptest.NewRequest(http.MethodPost, "/api/tags", bytes.NewBuffer([]byte("invalid json")))
	rec := httptest.NewRecorder()

	tagHandler.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	mockTagManager.AssertExpectations(t)
}

func TestTagHandler_Delete_Success(t *testing.T) {
	tagHandler, mockTagManager := setupTagHandlerTests()

	deletedTag := structs.Tag{Name: "TagToDelete"}
	mockTagManager.On("Delete", "TagToDelete").Return(deletedTag, nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/tags/TagToDelete", nil)
	req = mux.SetURLVars(req, map[string]string{"name": "TagToDelete"})
	rec := httptest.NewRecorder()

	tagHandler.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var result structs.Tag
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, deletedTag, result)

	mockTagManager.AssertExpectations(t)
}

func TestTagHandler_Delete_NotFound(t *testing.T) {
	tagHandler, mockTagManager := setupTagHandlerTests()

	mockTagManager.On("Delete", "NonexistentTag").
		Return(structs.Tag{}, errors.New("not found"))

	req := httptest.NewRequest(http.MethodDelete, "/api/tags/NonexistentTag", nil)
	req = mux.SetURLVars(req, map[string]string{"name": "NonexistentTag"})
	rec := httptest.NewRecorder()

	tagHandler.Delete(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	mockTagManager.AssertExpectations(t)
}
