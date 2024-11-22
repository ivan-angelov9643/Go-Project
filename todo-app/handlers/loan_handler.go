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

type LoanHandler struct {
	loanManager interfaces.LoanManager
}

func NewLoanHandler(loanManager interfaces.LoanManager) *LoanHandler {
	return &LoanHandler{loanManager}
}

func (h *LoanHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[LoanHandler.GetAll] Fetching all loans")

	loans := h.loanManager.GetAll()

	err := json.NewEncoder(w).Encode(loans)
	if err != nil {
		global.HttpError(
			w,
			"[LoanHandler.GetAll] Failed to encode loans to JSON",
			"Failed to return loans",
			http.StatusInternalServerError,
			err,
		)
		return
	}
}

func (h *LoanHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[LoanHandler.Get] Fetching loan")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(
			w,
			"[LoanHandler.Get] Invalid UUID format",
			"Invalid loan ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	loan, err := h.loanManager.Get(id)
	if err != nil {
		global.HttpError(
			w,
			"[LoanHandler.Get] Loan not found",
			"Loan not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(loan)
	if err != nil {
		global.HttpError(
			w,
			"[LoanHandler.Get] Failed to encode loan to JSON",
			"Failed to return loan",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *LoanHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("[LoanHandler.Create] Creating new loan")

	var newLoan models.Loan
	err := json.NewDecoder(r.Body).Decode(&newLoan)
	if err != nil {
		global.HttpError(
			w,
			"[LoanHandler.Create] Failed to decode JSON body into Loan struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newLoan.ID = uuid.Nil
	createdLoan, err := h.loanManager.Create(newLoan)
	if err != nil {
		global.HttpError(
			w,
			"[LoanHandler.Create] Error creating loan",
			"Failed to create loan",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdLoan)
	if err != nil {
		global.HttpError(w,
			"[LoanHandler.Create] Failed to encode created loan to JSON",
			"Failed to return created loan",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *LoanHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[LoanHandler.Update] Updating loan")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[LoanHandler.Update] Invalid UUID format",
			"Invalid loan ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedLoanBody models.Loan
	err = json.NewDecoder(r.Body).Decode(&updatedLoanBody)
	if err != nil {
		global.HttpError(
			w,
			"[LoanHandler.Update] Failed to decode JSON body into Loan struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedLoanBody.ID = id
	updatedLoan, err := h.loanManager.Update(updatedLoanBody)
	if err != nil {
		global.HttpError(w,
			"[LoanHandler.Update] Error updating loan",
			"Failed to update loan",
			http.StatusInternalServerError,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedLoan)
	if err != nil {
		global.HttpError(w,
			"[LoanHandler.Update] Failed to encode updated loan to JSON",
			"Failed to return updated loan",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *LoanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[LoanHandler.Delete] Deleting loan")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[LoanHandler.Delete] Invalid UUID format",
			"Invalid loan ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedLoan, err := h.loanManager.Delete(id)
	if err != nil {
		global.HttpError(w,
			"[LoanHandler.Delete] Loan not found",
			"Loan not found",
			http.StatusNotFound,
			err,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedLoan)
	if err != nil {
		global.HttpError(w,
			"[LoanHandler.Delete] Failed to encode loan to JSON",
			"Failed to return deleted loan",
			http.StatusInternalServerError,
			err,
		)
	}
}
