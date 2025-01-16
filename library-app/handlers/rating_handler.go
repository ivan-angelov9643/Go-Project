package handlers

import (
	"awesomeProject/library-app/global"
	"awesomeProject/library-app/global/db"
	"awesomeProject/library-app/managers"
	"awesomeProject/library-app/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type RatingHandler struct {
	ratingManager managers.RatingManagerInterface
}

func NewRatingHandler(ratingManager managers.RatingManagerInterface) *RatingHandler {
	return &RatingHandler{ratingManager}
}

func (h *RatingHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[RatingHandler.GetAll] Fetching all ratings")

	dbScope := db.NewDBScope(global.IsGlobal(r), global.GetOwnerID(r))
	ratings, dbErr := h.ratingManager.GetAll(dbScope)
	if dbErr != nil {
		global.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err := json.NewEncoder(w).Encode(ratings)
	if err != nil {
		global.HttpError(
			w,
			"[RatingHandler.GetAll] Failed to encode ratings to JSON",
			"Failed to return ratings",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *RatingHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[RatingHandler.Get] Fetching rating")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(
			w,
			"[RatingHandler.Get] Invalid UUID format",
			"Invalid rating ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	rating, dbErr := h.ratingManager.Get(id)
	if dbErr != nil {
		global.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(rating)
	if err != nil {
		global.HttpError(
			w,
			"[RatingHandler.Get] Failed to encode rating to JSON",
			"Failed to return rating",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *RatingHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[RatingHandler.Create] Creating new rating")

	var newRating models.Rating
	err := json.NewDecoder(r.Body).Decode(&newRating)
	if err != nil {
		global.HttpError(
			w,
			"[RatingHandler.Create] Failed to decode JSON body into Rating struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newRating.ID = uuid.Nil
	createdRating, dbErr := h.ratingManager.Create(newRating)
	if dbErr != nil {
		global.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdRating)
	if err != nil {
		global.HttpError(w,
			"[RatingHandler.Create] Failed to encode created rating to JSON",
			"Failed to return created rating",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *RatingHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[RatingHandler.Update] Updating rating")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[RatingHandler.Update] Invalid UUID format",
			"Invalid rating ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedRatingBody models.Rating
	err = json.NewDecoder(r.Body).Decode(&updatedRatingBody)
	if err != nil {
		global.HttpError(
			w,
			"[RatingHandler.Update] Failed to decode JSON body into Rating struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedRatingBody.ID = id
	updatedRating, dbErr := h.ratingManager.Update(updatedRatingBody)
	if dbErr != nil {
		global.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedRating)
	if err != nil {
		global.HttpError(w,
			"[RatingHandler.Update] Failed to encode updated rating to JSON",
			"Failed to return updated rating",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *RatingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[RatingHandler.Delete] Deleting rating")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[RatingHandler.Delete] Invalid UUID format",
			"Invalid rating ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedRating, dbErr := h.ratingManager.Delete(id)
	if dbErr != nil {
		global.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedRating)
	if err != nil {
		global.HttpError(w,
			"[RatingHandler.Delete] Failed to encode rating to JSON",
			"Failed to return deleted rating",
			http.StatusInternalServerError,
			err,
		)
	}
}
