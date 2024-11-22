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

type ReservationHandler struct {
	reservationManager interfaces.ReservationManager
}

func NewReservationHandler(reservationManager interfaces.ReservationManager) *ReservationHandler {
	return &ReservationHandler{reservationManager}
}

func (h *ReservationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReservationHandler.GetAll] Fetching all reservations")

	reservations := h.reservationManager.GetAll()

	err := json.NewEncoder(w).Encode(reservations)
	if err != nil {
		global.HttpError(
			w,
			"[ReservationHandler.GetAll] Failed to encode reservations to JSON",
			"Failed to return reservations",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *ReservationHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[ReservationHandler.Get] Fetching reservation")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(
			w,
			"[ReservationHandler.Get] Invalid UUID format",
			"Invalid reservation ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	reservation, err := h.reservationManager.Get(id)
	if err != nil {
		global.HttpError(
			w,
			"[ReservationHandler.Get] Reservation not found",
			"Reservation not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(reservation)
	if err != nil {
		global.HttpError(
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
		global.HttpError(
			w,
			"[ReservationHandler.Create] Failed to decode JSON body into Reservation struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newReservation.ID = uuid.Nil
	createdReservation, err := h.reservationManager.Create(newReservation)
	if err != nil {
		global.HttpError(
			w,
			"[ReservationHandler.Create] Error creating reservation",
			"Failed to create reservation",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdReservation)
	if err != nil {
		global.HttpError(w,
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
		global.HttpError(w,
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
		global.HttpError(
			w,
			"[ReservationHandler.Update] Failed to decode JSON body into Reservation struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedReservationBody.ID = id
	updatedReservation, err := h.reservationManager.Update(updatedReservationBody)
	if err != nil {
		global.HttpError(w,
			"[ReservationHandler.Update] Error updating reservation",
			"Failed to update reservation",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedReservation)
	if err != nil {
		global.HttpError(w,
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
		global.HttpError(w,
			"[ReservationHandler.Delete] Invalid UUID format",
			"Invalid reservation ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedReservation, err := h.reservationManager.Delete(id)
	if err != nil {
		global.HttpError(w,
			"[ReservationHandler.Delete] Reservation not found",
			"Reservation not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedReservation)
	if err != nil {
		global.HttpError(w,
			"[ReservationHandler.Delete] Failed to encode reservation to JSON",
			"Failed to return deleted reservation",
			http.StatusInternalServerError,
			err,
		)
	}
}