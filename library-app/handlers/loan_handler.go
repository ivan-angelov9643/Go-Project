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

type LoanHandler struct {
	loanManager managers.LoanManagerInterface
}

func NewLoanHandler(loanManager managers.LoanManagerInterface) *LoanHandler {
	return &LoanHandler{loanManager}
}

func (h *LoanHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[LoanHandler.GetAll] Fetching all loans")

	accessScope := db.NewAccessScope(r)
	pagingScope := db.NewPagingScope(r)
	filterByBookIDScope := db.NewFilterByBookIDScope(r)
	filterByUserIDScope := db.NewFilterByUserIDScope(r)
	filterByUsernameScope := db.NewFilterByUsernameScope(r)
	filterByTitleScope := db.NewFilterByTitleScope(r)
	filterByStatusScope := db.NewFilterByStatusScope(r)
	sortScope := db.NewSortScope(r)
	loans, dbErr := h.loanManager.GetAll(
		accessScope, pagingScope, filterByBookIDScope, filterByUserIDScope, filterByUsernameScope, filterByTitleScope,
		filterByStatusScope, sortScope,
	)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	count, dbErr := h.loanManager.Count(
		accessScope, filterByBookIDScope, filterByUserIDScope, filterByUsernameScope, filterByTitleScope,
		filterByStatusScope,
	)
	if dbErr != nil {
		errors.HttpDBError(w, dbErr)
		return
	}

	response := global.PaginatedResponse[models.Loan]{
		Count:    count,
		PageSize: pagingScope.PageSize,
		Page:     pagingScope.Page,
		Data:     loans,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errors.HttpError(
			w,
			"[LoanHandler.GetAll] Failed to encode loans to JSON",
			"Failed to return loans",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *LoanHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[LoanHandler.Get] Fetching loan")

	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		errors.HttpError(
			w,
			"[LoanHandler.Get] Invalid UUID format",
			"Invalid loan ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	loan, dbErr := h.loanManager.Get(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(loan)
	if err != nil {
		errors.HttpError(
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
		errors.HttpError(
			w,
			"[LoanHandler.Create] Failed to decode JSON body into Loan struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	newLoan.ID = uuid.Nil
	createdLoan, dbErr := h.loanManager.Create(newLoan)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(createdLoan)
	if err != nil {
		errors.HttpError(w,
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
		errors.HttpError(w,
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
		errors.HttpError(
			w,
			"[LoanHandler.Update] Failed to decode JSON body into Loan struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedLoanBody.ID = id
	updatedLoan, dbErr := h.loanManager.Update(updatedLoanBody)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(updatedLoan)
	if err != nil {
		errors.HttpError(w,
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
		errors.HttpError(w,
			"[LoanHandler.Delete] Invalid UUID format",
			"Invalid loan ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedLoan, dbErr := h.loanManager.Delete(id)
	if dbErr != nil {
		errors.HttpDBError(
			w,
			dbErr,
		)
		return
	}

	err = json.NewEncoder(w).Encode(deletedLoan)
	if err != nil {
		errors.HttpError(w,
			"[LoanHandler.Delete] Failed to encode loan to JSON",
			"Failed to return deleted loan",
			http.StatusInternalServerError,
			err,
		)
	}
}
