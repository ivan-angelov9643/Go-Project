package db

import (
	"gorm.io/gorm"
	"net/http"
)

type FilterByAuthorNameScope struct {
	Search *string
}

func NewFilterByAuthorNameScope(r *http.Request) *FilterByAuthorNameScope {
	searchParam := r.URL.Query().Get("search")

	var search *string = nil
	if searchParam != "" && searchParam != "null" {
		search = &searchParam
	}

	return &FilterByAuthorNameScope{Search: search}
}

func (s *FilterByAuthorNameScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Search == nil {
			return db
		} else {
			searchTerm := "%" + *s.Search + "%"
			return db.Where("CONCAT(authors.first_name, ' ', authors.last_name) ILIKE ?", searchTerm)
		}
	}
}
