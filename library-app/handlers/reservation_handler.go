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

type ReservationHandler struct {
	reservationManager managers.ReservationManagerInterface
}

func NewReservationHandler(reservationManager managers.ReservationManagerInterface) *ReservationHandler {
	return &ReservationHandler{reservationManager}
}

func (h *ReservationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReservationHandler.GetAll] Fetching all reservations")

	accessScope := db.NewAccessScope(r)
	pagingScope := db.NewPagingScope(r)
	filterByBookIDScope := db.NewFilterByBookIDScope(r)
	filterByUserIDScope := db.NewFilterByUserIDScope(r)
	filterByUsernameScope := db.NewFilterByUsernameScope(r)
	filterByTitleScope := db.NewFilterByTitleScope(r)
	sortScope := db.NewSortScope(r)
	reservations, dbErr := h.reservationManager.GetAll(
		accessScope, pagingScope, filterByUserIDScope, filterByBookIDScope, filterByUsernameScope, filterByTitleScope,
		sortScope,
	)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	count, dbErr := h.reservationManager.Count(
		accessScope, filterByUserIDScope, filterByBookIDScope, filterByUsernameScope, filterByTitleScope,
	)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	response := global.PaginatedResponse[models.Reservation]{
		Count:    count,
		PageSize: pagingScope.PageSize,
		Page:     pagingScope.Page,
		Data:     reservations,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errors.HttpError(
			w,
			"[ReservationHandler.GetAll] Failed to encode reservations to JSON",
			"Failed to return reservations",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ReservationHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReservationHandler.Get] Fetching reservation")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(
			w,
			"[ReservationHandler.Get] Invalid UUID format",
			"Invalid reservation ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	reservation, dbErr := h.reservationManager.Get(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(reservation)
	if err != nil {
		errors.HttpError(
			w,
			"[ReservationHandler.Get] Failed to encode reservation to JSON",
			"Failed to return reservation",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReservationHandler.Create] Creating new reservation")

	var newReservation models.Reservation
	err := json.NewDecoder(r.Body).Decode(&newReservation)
	if err != nil {
		errors.HttpError(
			w,
			"[ReservationHandler.Create] Failed to decode JSON body into Reservation struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newReservation.ID = uuid.Nil
	createdReservation, dbErr := h.reservationManager.Create(newReservation)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdReservation)
	if err != nil {
		errors.HttpError(w,
			"[ReservationHandler.Create] Failed to encode created reservation to JSON",
			"Failed to return created reservation",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ReservationHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReservationHandler.Update] Updating reservation")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(w,
			"[ReservationHandler.Update] Invalid UUID format",
			"Invalid reservation ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedReservationBody models.Reservation
	err = json.NewDecoder(r.Body).Decode(&updatedReservationBody)
	if err != nil {
		errors.HttpError(
			w,
			"[ReservationHandler.Update] Failed to decode JSON body into Reservation struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedReservationBody.ID = id
	updatedReservation, dbErr := h.reservationManager.Update(updatedReservationBody)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedReservation)
	if err != nil {
		errors.HttpError(w,
			"[ReservationHandler.Update] Failed to encode updated reservation to JSON",
			"Failed to return updated reservation",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *ReservationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReservationHandler.Delete] Deleting reservation")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(w,
			"[ReservationHandler.Delete] Invalid UUID format",
			"Invalid reservation ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedReservation, dbErr := h.reservationManager.Delete(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedReservation)
	if err != nil {
		errors.HttpError(w,
			"[ReservationHandler.Delete] Failed to encode reservation to JSON",
			"Failed to return deleted reservation",
			http.StatusInternalServerError,
			err,
		)
	}
}
