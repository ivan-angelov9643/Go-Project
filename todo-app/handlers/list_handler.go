package handlers

import (
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/managers/interfaces"
	"awesomeProject/todo-app/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ListHandler struct {
	listManager interfaces.ListManager
}

func NewListHandler(listManager interfaces.ListManager) *ListHandler {
	return &ListHandler{listManager}
}

func (h *ListHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[ListHandler.GetAll] Fetching all lists")

	lists := h.listManager.GetAll()

	err := json.NewEncoder(w).Encode(lists)
	if err != nil {
		global.HttpError(
			w,
			"[ListHandler.GetAll] Failed to encode lists to JSON",
			"Failed to return lists",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *ListHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[ListHandler.Get] Fetching list")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(
			w,
			"[ListHandler.Get] Invalid UUID format",
			"Invalid list ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	list, err := h.listManager.Get(id)
	if err != nil {
		global.HttpError(
			w,
			"[ListHandler.Get] List not found",
			"List not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		global.HttpError(
			w,
			"[ListHandler.Get] Failed to encode list to JSON",
			"Failed to return list",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ListHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[ListHandler.Create] Creating new list")

	var newList models.List
	err := json.NewDecoder(r.Body).Decode(&newList)
	if err != nil {
		global.HttpError(
			w,
			"[ListHandler.Create] Failed to decode JSON body into List struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newList.ID = uuid.Nil
	createdList, err := h.listManager.Create(newList)
	if err != nil {
		global.HttpError(
			w,
			"[ListHandler.Create] Error creating list",
			"Failed to create list",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdList)
	if err != nil {
		global.HttpError(w,
			"[ListHandler.Create] Failed to encode created list to JSON",
			"Failed to return created list",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ListHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[ListHandler.Update] Updating list")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[ListHandler.Update] Invalid UUID format",
			"Invalid list ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedListBody models.List
	err = json.NewDecoder(r.Body).Decode(&updatedListBody)
	if err != nil {
		global.HttpError(
			w,
			"[ListHandler.Update] Failed to decode JSON body into List struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedListBody.ID = id
	updatedList, err := h.listManager.Update(updatedListBody)
	if err != nil {
		global.HttpError(w,
			"[ListHandler.Update] Error updating list",
			"Failed to update list",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedList)
	if err != nil {
		global.HttpError(w,
			"[ListHandler.Update] Failed to encode updated list to JSON",
			"Failed to return updated list",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[ListHandler.Delete] Deleting list")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[ListHandler.Delete] Invalid UUID format",
			"Invalid list ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedList, err := h.listManager.Delete(id)
	if err != nil {
		global.HttpError(w,
			"[ListHandler.Delete] List not found",
			"List not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedList)
	if err != nil {
		global.HttpError(w,
			"[ListHandler.Delete] Failed to encode list to JSON",
			"Failed to return deleted list",
			http.StatusInternalServerError,
			err,
		)
	}
}
