package db

import (
	"gorm.io/gorm"
	"net/http"
)

type FilterByUsernameScope struct {
	Search *string
}

func NewFilterByUsernameScope(r *http.Request) *FilterByUsernameScope {
	searchParam := r.URL.Query().Get("username")

	var search *string = nil
	if searchParam != "" && searchParam != "null" {
		search = &searchParam
	}

	return &FilterByUsernameScope{Search: search}
}

func (s *FilterByUsernameScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Search == nil {
			return db
		} else {
			searchTerm := "%" + *s.Search + "%"
			return db.Where("users.preferred_username ILIKE ?", searchTerm)
		}
	}
}
