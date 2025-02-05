package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type FilterByUserIDScope struct {
	UserID *uuid.UUID
}

func NewFilterByUserIDScope(r *http.Request) *FilterByUserIDScope {
	userIDParam := r.URL.Query().Get("user_id")

	var userID *uuid.UUID = nil
	if userIDParam != "" && userIDParam != "null" {
		parsedUserID, err := uuid.Parse(userIDParam)
		if err == nil {
			userID = &parsedUserID
		}
	}
	return &FilterByUserIDScope{userID}
}

func (s *FilterByUserIDScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.UserID == nil {
			return db
		} else {
			return db.Where("user_id = ?", s.UserID)
		}
	}
}
