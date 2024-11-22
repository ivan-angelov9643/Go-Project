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

type BookHandler struct {
	bookManager interfaces.BookManager
}

func NewBookHandler(bookManager interfaces.BookManager) *BookHandler {
	return &BookHandler{bookManager}
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[BookHandler.GetAll] Fetching all books")

	books := h.bookManager.GetAll()

	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		global.HttpError(
			w,
			"[BookHandler.GetAll] Failed to encode books to JSON",
			"Failed to return books",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[BookHandler.Get] Fetching book")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(
			w,
			"[BookHandler.Get] Invalid UUID format",
			"Invalid book ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	book, err := h.bookManager.Get(id)
	if err != nil {
		global.HttpError(
			w,
			"[BookHandler.Get] Book not found",
			"Book not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		global.HttpError(
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
		global.HttpError(
			w,
			"[BookHandler.Create] Failed to decode JSON body into Book struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newBook.ID = uuid.Nil
	createdBook, err := h.bookManager.Create(newBook)
	if err != nil {
		global.HttpError(
			w,
			"[BookHandler.Create] Error creating book",
			"Failed to create book",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdBook)
	if err != nil {
		global.HttpError(w,
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
		global.HttpError(w,
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
		global.HttpError(
			w,
			"[BookHandler.Update] Failed to decode JSON body into Book struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedBookBody.ID = id
	updatedBook, err := h.bookManager.Update(updatedBookBody)
	if err != nil {
		global.HttpError(w,
			"[BookHandler.Update] Error updating book",
			"Failed to update book",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		global.HttpError(w,
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
		global.HttpError(w,
			"[BookHandler.Delete] Invalid UUID format",
			"Invalid book ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedBook, err := h.bookManager.Delete(id)
	if err != nil {
		global.HttpError(w,
			"[BookHandler.Delete] Book not found",
			"Book not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedBook)
	if err != nil {
		global.HttpError(w,
			"[BookHandler.Delete] Failed to encode book to JSON",
			"Failed to return deleted book",
			http.StatusInternalServerError,
			err,
		)
	}
}
