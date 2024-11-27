package handlers

import (
	"awesomeProject/library-app/global"
	"awesomeProject/library-app/managers/interfaces"
	"awesomeProject/library-app/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ItemHandler struct {
	itemManager interfaces.ItemManager
}

func NewItemHandler(itemManager interfaces.ItemManager) *ItemHandler {
	return &ItemHandler{itemManager}
}

func (h *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[ItemHandler.GetAll] Fetching all items")

	items := h.itemManager.GetAll()

	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		global.HttpError(
			w,
			"[ItemHandler.GetAll] Failed to encode items to JSON",
			"Failed to return items",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[ItemHandler.Get] Fetching item")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(
			w,
			"[ItemHandler.Get] Invalid UUID format",
			"Invalid item ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	item, err := h.itemManager.Get(id)
	if err != nil {
		global.HttpError(
			w,
			"[ItemHandler.Get] Item not found",
			"Item not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		global.HttpError(
			w,
			"[ItemHandler.Get] Failed to encode item to JSON",
			"Failed to return item",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[ItemHandler.Create] Creating new item")

	var newItem models.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		global.HttpError(
			w,
			"[ItemHandler.Create] Failed to decode JSON body into Item struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newItem.ID = uuid.Nil
	createdItem, err := h.itemManager.Create(newItem)
	if err != nil {
		global.HttpError(
			w,
			"[ItemHandler.Create] Error creating item",
			"Failed to create item",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdItem)
	if err != nil {
		global.HttpError(w,
			"[ItemHandler.Create] Failed to encode created item to JSON",
			"Failed to return created item",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[ItemHandler.Update] Updating item")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[ItemHandler.Update] Invalid UUID format",
			"Invalid item ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedItemBody models.Item
	err = json.NewDecoder(r.Body).Decode(&updatedItemBody)
	if err != nil {
		global.HttpError(
			w,
			"[ItemHandler.Update] Failed to decode JSON body into Item struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedItemBody.ID = id
	updatedItem, err := h.itemManager.Update(updatedItemBody)
	if err != nil {
		global.HttpError(w,
			"[ItemHandler.Update] Error updating item",
			"Failed to update item",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedItem)
	if err != nil {
		global.HttpError(w,
			"[ItemHandler.Update] Failed to encode updated item to JSON",
			"Failed to return updated item",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[ItemHandler.Delete] Deleting item")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[ItemHandler.Delete] Invalid UUID format",
			"Invalid item ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedItem, err := h.itemManager.Delete(id)
	if err != nil {
		global.HttpError(w,
			"[ItemHandler.Delete] Item not found",
			"Item not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedItem)
	if err != nil {
		global.HttpError(w,
			"[ItemHandler.Delete] Failed to encode item to JSON",
			"Failed to return deleted item",
			http.StatusInternalServerError,
			err,
		)
	}
}
