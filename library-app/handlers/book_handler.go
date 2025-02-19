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

type BookHandler struct {
	bookManager managers.BookManagerInterface
}

func NewBookHandler(bookManager managers.BookManagerInterface) *BookHandler {
	return &BookHandler{bookManager}
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[BookHandler.GetAll] Fetching all books")

	accessScope := db.NewAccessScope(r)
	pagingScope := db.NewPagingScope(r)
	filterByTitleScope := db.NewFilterByTitleScope(r)
	filterByAuthorNameScope := db.NewFilterByAuthorNameScope(r)
	filterByCategoryNameScope := db.NewFilterByCategoryNameScope(r)
	filterByLanguageScope := db.NewFilterByLanguageScope(r)
	books, dbErr := h.bookManager.GetAll(
		accessScope, pagingScope, filterByTitleScope, filterByAuthorNameScope, filterByCategoryNameScope, filterByLanguageScope,
	)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	count, dbErr := h.bookManager.Count(
		accessScope, filterByTitleScope, filterByAuthorNameScope, filterByCategoryNameScope, filterByLanguageScope,
	)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	response := global.PaginatedResponse[models.Book]{
		Count:    count,
		PageSize: pagingScope.PageSize,
		Page:     pagingScope.Page,
		Data:     books,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errors.HttpError(
			w,
			"[BookHandler.GetAll] Failed to encode books to JSON",
			"Failed to return books",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[BookHandler.Get] Fetching book")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(
			w,
			"[BookHandler.Get] Invalid UUID format",
			"Invalid book ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	book, dbErr := h.bookManager.Get(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		errors.HttpError(
			w,
			"[BookHandler.Get] Failed to encode book to JSON",
			"Failed to return book",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[BookHandler.Create] Creating new book")

	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		errors.HttpError(
			w,
			"[BookHandler.Create] Failed to decode JSON body into Book struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newBook.ID = uuid.Nil
	createdBook, dbErr := h.bookManager.Create(newBook)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdBook)
	if err != nil {
		errors.HttpError(w,
			"[BookHandler.Create] Failed to encode created book to JSON",
			"Failed to return created book",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[BookHandler.Update] Updating book")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(w,
			"[BookHandler.Update] Invalid UUID format",
			"Invalid book ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedBookBody models.Book
	err = json.NewDecoder(r.Body).Decode(&updatedBookBody)
	if err != nil {
		errors.HttpError(
			w,
			"[BookHandler.Update] Failed to decode JSON body into Book struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedBookBody.ID = id
	updatedBook, dbErr := h.bookManager.Update(updatedBookBody)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		errors.HttpError(w,
			"[BookHandler.Update] Failed to encode updated book to JSON",
			"Failed to return updated book",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[BookHandler.Delete] Deleting book")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(w,
			"[BookHandler.Delete] Invalid UUID format",
			"Invalid book ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedBook, dbErr := h.bookManager.Delete(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedBook)
	if err != nil {
		errors.HttpError(w,
			"[BookHandler.Delete] Failed to encode book to JSON",
			"Failed to return deleted book",
			http.StatusInternalServerError,
			err,
		)
	}
}
