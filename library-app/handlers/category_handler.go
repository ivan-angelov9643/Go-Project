package handlers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/errors"
	"awesomeProject/library-app/global"
	"awesomeProject/library-app/managers"
	"awesomeProject/library-app/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type CategoryHandler struct {
	categoryManager managers.CategoryManagerInterface
}

func NewCategoryHandler(categoryManager managers.CategoryManagerInterface) *CategoryHandler {
	return &CategoryHandler{categoryManager}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[CategoryHandler.GetAll] Fetching all categories")

	accessScope := db.NewAccessScope(r)
	pagingScope := db.NewPagingScope(r)
	filterByCategoryNameScope := db.NewFilterByCategoryNameScope(r)
	categories, dbErr := h.categoryManager.GetAll(accessScope, pagingScope, filterByCategoryNameScope)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	count, dbErr := h.categoryManager.Count(accessScope, filterByCategoryNameScope)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	response := global.PaginatedResponse[models.Category]{
		Count:    count,
		PageSize: pagingScope.PageSize,
		Page:     pagingScope.Page,
		Data:     categories,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errors.HttpError(
			w,
			"[CategoryHandler.GetAll] Failed to encode categories to JSON",
			"Failed to return categories",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *CategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[CategoryHandler.Get] Fetching category")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(
			w,
			"[CategoryHandler.Get] Invalid UUID format",
			"Invalid category ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	category, dbErr := h.categoryManager.Get(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		errors.HttpError(
			w,
			"[CategoryHandler.Get] Failed to encode category to JSON",
			"Failed to return category",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[CategoryHandler.Create] Creating new category")

	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		errors.HttpError(
			w,
			"[CategoryHandler.Create] Failed to decode JSON body into Category struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newCategory.ID = uuid.Nil
	createdCategory, dbErr := h.categoryManager.Create(newCategory)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdCategory)
	if err != nil {
		errors.HttpError(w,
			"[CategoryHandler.Create] Failed to encode created category to JSON",
			"Failed to return created category",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[CategoryHandler.Update] Updating category")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(w,
			"[CategoryHandler.Update] Invalid UUID format",
			"Invalid category ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedCategoryBody models.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategoryBody)
	if err != nil {
		errors.HttpError(
			w,
			"[CategoryHandler.Update] Failed to decode JSON body into Category struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedCategoryBody.ID = id
	updatedCategory, dbErr := h.categoryManager.Update(updatedCategoryBody)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedCategory)
	if err != nil {
		errors.HttpError(w,
			"[CategoryHandler.Update] Failed to encode updated category to JSON",
			"Failed to return updated category",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[CategoryHandler.Delete] Deleting category")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(w,
			"[CategoryHandler.Delete] Invalid UUID format",
			"Invalid category ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedCategory, dbErr := h.categoryManager.Delete(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedCategory)
	if err != nil {
		errors.HttpError(w,
			"[CategoryHandler.Delete] Failed to encode category to JSON",
			"Failed to return deleted category",
			http.StatusInternalServerError,
			err,
		)
	}
}
