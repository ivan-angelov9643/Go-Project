package handlers

import (
	"awesomeProject/library-app/global"
	"awesomeProject/library-app/managers"
	"awesomeProject/library-app/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler struct {
	userManager managers.UserManagerInterface
}

func NewUserHandler(userManager managers.UserManagerInterface) *UserHandler {
	return &UserHandler{userManager}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[UserHandler.GetAll] Fetching all users")

	users, dbErr := h.userManager.GetAll()
	if dbErr != nil {
		global.HttpDBError(w, dbErr)
		return
	}

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		global.HttpError(w,
			"[UserHandler.GetAll] Failed to encode users to JSON",
			"Failed to return users",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("[UserHandler.Get] Fetching user")

	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[UserHandler.Get] Invalid UUID format",
			"Invalid user ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	user, dbErr := h.userManager.Get(id)
	if dbErr != nil {
		//global.HttpDBError(w, dbErr)/
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		global.HttpError(w,
			"[UserHandler.Get] Failed to encode user to JSON",
			"Failed to return user",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("[UserHandler.Update] Updating user")

	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[UserHandler.Update] Invalid UUID format",
			"Invalid user ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	var updatedUserBody models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUserBody)
	if err != nil {
		global.HttpError(w,
			"[UserHandler.Update] Failed to decode JSON body into User struct",
			"Invalid JSON format in request body",
			http.StatusBadRequest,
			err,
		)
		return
	}

	updatedUserBody.ID = id
	updatedUser, dbErr := h.userManager.Update(updatedUserBody)
	if dbErr != nil {
		global.HttpDBError(w, dbErr)
		return
	}

	err = json.NewEncoder(w).Encode(updatedUser)
	if err != nil {
		global.HttpError(w,
			"[UserHandler.Update] Failed to encode updated user to JSON",
			"Failed to return updated user",
			http.StatusInternalServerError,
			err,
		)
	}
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("[UserHandler.Delete] Deleting user")

	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		global.HttpError(w,
			"[UserHandler.Delete] Invalid UUID format",
			"Invalid user ID format",
			http.StatusBadRequest,
			err,
		)
		return
	}

	deletedUser, dbErr := h.userManager.Delete(id)
	if dbErr != nil {
		global.HttpDBError(w, dbErr)
		return
	}

	err = json.NewEncoder(w).Encode(deletedUser)
	if err != nil {
		global.HttpError(w,
			"[UserHandler.Delete] Failed to encode user to JSON",
			"Failed to return deleted user",
			http.StatusInternalServerError,
			err,
		)
	}
}
