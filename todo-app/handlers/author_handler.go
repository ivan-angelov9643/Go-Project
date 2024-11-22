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

type AuthorHandler struct {
	authorManager interfaces.AuthorManager
}

func NewAuthorHandler(authorManager interfaces.AuthorManager) *AuthorHandler {
	return &AuthorHandler{authorManager}
}

func (h *AuthorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[AuthorHandler.GetAll] Fetching all authors")

	authors := h.authorManager.GetAll()

	err := json.NewEncoder(w).Encode(authors)
	if err != nil {
		global.HttpError(
			w,
			"[AuthorHandler.GetAll] Failed to encode authors to JSON",
			"Failed to return authors",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *AuthorHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[AuthorHandler.Get] Fetching author")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(
			w,
			"[AuthorHandler.Get] Invalid UUID format",
			"Invalid author ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	author, err := h.authorManager.Get(id)
	if err != nil {
		global.HttpError(
			w,
			"[AuthorHandler.Get] Author not found",
			"Author not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(author)
	if err != nil {
		global.HttpError(
			w,
			"[AuthorHandler.Get] Failed to encode author to JSON",
			"Failed to return author",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *AuthorHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[AuthorHandler.Create] Creating new author")

	var newAuthor models.Author
	err := json.NewDecoder(r.Body).Decode(&newAuthor)
	if err != nil {
		global.HttpError(
			w,
			"[AuthorHandler.Create] Failed to decode JSON body into Author struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newAuthor.ID = uuid.Nil
	createdAuthor, err := h.authorManager.Create(newAuthor)
	if err != nil {
		global.HttpError(
			w,
			"[AuthorHandler.Create] Error creating author",
			"Failed to create author",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdAuthor)
	if err != nil {
		global.HttpError(w,
			"[AuthorHandler.Create] Failed to encode created author to JSON",
			"Failed to return created author",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *AuthorHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[AuthorHandler.Update] Updating author")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[AuthorHandler.Update] Invalid UUID format",
			"Invalid author ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedAuthorBody models.Author
	err = json.NewDecoder(r.Body).Decode(&updatedAuthorBody)
	if err != nil {
		global.HttpError(
			w,
			"[AuthorHandler.Update] Failed to decode JSON body into Author struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	//err := updatedAuthor.Validate()
	//if err != nil {
	//	return models.Author{}, err
	//}
	updatedAuthorBody.ID = id
	updatedAuthor, err := h.authorManager.Update(updatedAuthorBody)
	if err != nil {
		global.HttpError(w,
			"[AuthorHandler.Update] Error updating author",
			"Failed to update author",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedAuthor)
	if err != nil {
		global.HttpError(w,
			"[AuthorHandler.Update] Failed to encode updated author to JSON",
			"Failed to return updated author",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *AuthorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[AuthorHandler.Delete] Deleting author")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[AuthorHandler.Delete] Invalid UUID format",
			"Invalid author ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedAuthor, err := h.authorManager.Delete(id)
	if err != nil {
		global.HttpError(w,
			"[AuthorHandler.Delete] Author not found",
			"Author not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedAuthor)
	if err != nil {
		global.HttpError(w,
			"[AuthorHandler.Delete] Failed to encode author to JSON",
			"Failed to return deleted author",
			http.StatusInternalServerError,
			err,
		)
	}
}
