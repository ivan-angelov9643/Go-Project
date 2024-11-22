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

type ReviewHandler struct {
	reviewManager interfaces.ReviewManager
}

func NewReviewHandler(reviewManager interfaces.ReviewManager) *ReviewHandler {
	return &ReviewHandler{reviewManager}
}

func (h *ReviewHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReviewHandler.GetAll] Fetching all reviews")

	reviews := h.reviewManager.GetAll()

	err := json.NewEncoder(w).Encode(reviews)
	if err != nil {
		global.HttpError(
			w,
			"[ReviewHandler.GetAll] Failed to encode reviews to JSON",
			"Failed to return reviews",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *ReviewHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReviewHandler.Get] Fetching review")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(
			w,
			"[ReviewHandler.Get] Invalid UUID format",
			"Invalid review ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	review, err := h.reviewManager.Get(id)
	if err != nil {
		global.HttpError(
			w,
			"[ReviewHandler.Get] Review not found",
			"Review not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(review)
	if err != nil {
		global.HttpError(
			w,
			"[ReviewHandler.Get] Failed to encode review to JSON",
			"Failed to return review",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReviewHandler.Create] Creating new review")

	var newReview models.Review
	err := json.NewDecoder(r.Body).Decode(&newReview)
	if err != nil {
		global.HttpError(
			w,
			"[ReviewHandler.Create] Failed to decode JSON body into Review struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newReview.ID = uuid.Nil
	createdReview, err := h.reviewManager.Create(newReview)
	if err != nil {
		global.HttpError(
			w,
			"[ReviewHandler.Create] Error creating review",
			"Failed to create review",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdReview)
	if err != nil {
		global.HttpError(w,
			"[ReviewHandler.Create] Failed to encode created review to JSON",
			"Failed to return created review",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ReviewHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReviewHandler.Update] Updating review")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[ReviewHandler.Update] Invalid UUID format",
			"Invalid review ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedReviewBody models.Review
	err = json.NewDecoder(r.Body).Decode(&updatedReviewBody)
	if err != nil {
		global.HttpError(
			w,
			"[ReviewHandler.Update] Failed to decode JSON body into Review struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedReviewBody.ID = id
	updatedReview, err := h.reviewManager.Update(updatedReviewBody)
	if err != nil {
		global.HttpError(w,
			"[ReviewHandler.Update] Error updating review",
			"Failed to update review",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedReview)
	if err != nil {
		global.HttpError(w,
			"[ReviewHandler.Update] Failed to encode updated review to JSON",
			"Failed to return updated review",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ReviewHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReviewHandler.Delete] Deleting review")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[ReviewHandler.Delete] Invalid UUID format",
			"Invalid review ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedReview, err := h.reviewManager.Delete(id)
	if err != nil {
		global.HttpError(w,
			"[ReviewHandler.Delete] Review not found",
			"Review not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedReview)
	if err != nil {
		global.HttpError(w,
			"[ReviewHandler.Delete] Failed to encode review to JSON",
			"Failed to return deleted review",
			http.StatusInternalServerError,
			err,
		)
	}
}
