package handlers

import (
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/managers"
	"awesomeProject/todo-app/structs"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type TagHandler struct {
	managers.TagManager
}

func NewTagHandler() *TagHandler {
	return &TagHandler{*managers.NewTagManager()}
}

func (h *TagHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.GetAll] Fetching all tags")

	tags := h.TagManager.GetAll()

	err := json.NewEncoder(w).Encode(tags)
	if err != nil {
		global.HttpError(
			w,
			"[TagHandler.GetAll] Failed to encode tags to JSON",
			"Failed to retrieve tags",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *TagHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.Get] Fetching tag")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Get] Invalid UUID format",
			"Invalid tag ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	tag, err := h.TagManager.Get(id)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Get] Tag not found",
			"Tag not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Get] Failed to encode tag to JSON",
			"Failed to retrieve tag",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *TagHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.Create] Creating new tag")

	var newTag structs.Tag
	err := json.NewDecoder(r.Body).Decode(&newTag)
	if err != nil {
		global.HttpError(
			w,
			"[TagHandler.Create] Failed to decode JSON body into Tag struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	_, err = h.TagManager.Create(newTag)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Create] Error creating tag",
			"Failed to create tag",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(newTag)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Create] Failed to encode created tag to JSON",
			"Failed to return created tag",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *TagHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.Update] Updating tag")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Update] Invalid UUID format",
			"Invalid tag ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedTag structs.Tag
	updatedTag.ID = id
	err = json.NewDecoder(r.Body).Decode(&updatedTag)
	if err != nil {
		global.HttpError(
			w,
			"[TagHandler.Update] Failed to decode JSON body into Tag struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	_, err = h.TagManager.Update(updatedTag)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Update] Error updating tag",
			"Failed to update tag",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedTag)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Update] Failed to encode updated tag to JSON",
			"Failed to return updated tag",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *TagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.Delete] Deleting tag")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Delete] Invalid UUID format",
			"Invalid tag ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	tag, err := h.TagManager.Delete(id)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Delete] Tag not found",
			"Tag not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Delete] Failed to encode tag to JSON",
			"Failed to retrieve tag",
			http.StatusInternalServerError,
			err,
		)
	}
}
