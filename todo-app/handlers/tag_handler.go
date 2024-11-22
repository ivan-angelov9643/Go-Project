package handlers

import (
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/managers/interfaces"
	"awesomeProject/todo-app/models"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type TagHandler struct {
	tagManager interfaces.TagManager
}

func NewTagHandler(tagManager interfaces.TagManager) *TagHandler {
	return &TagHandler{tagManager}
}

func (h *TagHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.GetAll] Fetching all tags")

	tags := h.tagManager.GetAll()

	err := json.NewEncoder(w).Encode(tags)
	if err != nil {
		global.HttpError(
			w,
			"[TagHandler.GetAll] Failed to encode tags to JSON",
			"Failed to return tags",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *TagHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.Get] Fetching tag")

	vars := mux.Vars(r)

	name := vars["name"]

	tag, err := h.tagManager.Get(name)
	if err != nil {
		global.HttpError(
			w,
			"[TagHandler.Get] Tag not found",
			"Tag not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		global.HttpError(
			w,
			"[TagHandler.Get] Failed to encode tag to JSON",
			"Failed to return tag",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *TagHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.Create] Creating new tag")

	var newTag models.Tag
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

	createdTag, err := h.tagManager.Create(newTag)
	if err != nil {
		global.HttpError(
			w,
			"[TagHandler.Create] Error creating tag",
			"Failed to create tag",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdTag)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Create] Failed to encode created tag to JSON",
			"Failed to return created tag",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *TagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[TagHandler.Delete] Deleting tag")

	vars := mux.Vars(r)

	name := vars["name"]

	deletedTag, err := h.tagManager.Delete(name)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Delete] Tag not found",
			"Tag not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedTag)
	if err != nil {
		global.HttpError(w,
			"[TagHandler.Delete] Failed to encode tag to JSON",
			"Failed to return deleted tag",
			http.StatusInternalServerError,
			err,
		)
	}
}
