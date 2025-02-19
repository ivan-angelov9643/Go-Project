package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/errors"
	"github.com/ivan-angelov9643/go-project/library-app/global"
	"github.com/ivan-angelov9643/go-project/library-app/managers"
	"github.com/ivan-angelov9643/go-project/library-app/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type AuthorHandler struct {
	authorManager managers.AuthorManagerInterface
}

func NewAuthorHandler(authorManager managers.AuthorManagerInterface) *AuthorHandler {
	return &AuthorHandler{authorManager}
}

func (h *AuthorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[AuthorHandler.GetAll] Fetching all authors")

	accessScope := db.NewAccessScope(r)
	pagingScope := db.NewPagingScope(r)
	filterByAuthorNameScope := db.NewFilterByAuthorNameScope(r)
	authors, dbErr := h.authorManager.GetAll(accessScope, pagingScope, filterByAuthorNameScope)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	count, dbErr := h.authorManager.Count(accessScope, filterByAuthorNameScope)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	response := global.PaginatedResponse[models.Author]{
		Count:    count,
		PageSize: pagingScope.PageSize,
		Page:     pagingScope.Page,
		Data:     authors,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		errors.HttpError(
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
		errors.HttpError(
			w,
			"[AuthorHandler.Get] Invalid UUID format",
			"Invalid author ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	author, dbErr := h.authorManager.Get(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(author)
	if err != nil {
		errors.HttpError(
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
		errors.HttpError(
			w,
			"[AuthorHandler.Create] Failed to decode JSON body into Author struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newAuthor.ID = uuid.Nil
	createdAuthor, dbErr := h.authorManager.Create(newAuthor)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	err = json.NewEncoder(w).Encode(createdAuthor)
	if err != nil {
		errors.HttpError(w,
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
		errors.HttpError(w,
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
		errors.HttpError(
			w,
			"[AuthorHandler.Update] Failed to decode JSON body into Author struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedAuthorBody.ID = id
	updatedAuthor, dbErr := h.authorManager.Update(updatedAuthorBody)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	err = json.NewEncoder(w).Encode(updatedAuthor)
	if err != nil {
		errors.HttpError(w,
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
		errors.HttpError(w,
			"[AuthorHandler.Delete] Invalid UUID format",
			"Invalid author ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedAuthor, dbErr := h.authorManager.Delete(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedAuthor)
	if err != nil {
		errors.HttpError(w,
			"[AuthorHandler.Delete] Failed to encode author to JSON",
			"Failed to return deleted author",
			http.StatusInternalServerError,
			err,
		)
	}
}
