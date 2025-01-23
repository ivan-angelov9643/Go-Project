package db

import (
	"gorm.io/gorm"
	"net/http"
)

type FilterByLanguageScope struct {
	Search *string
}

func NewFilterByLanguageScope(r *http.Request) *FilterByLanguageScope {
	searchParam := r.URL.Query().Get("language")

	var search *string = nil
	if searchParam != "" && searchParam != "null" {
		search = &searchParam
	}

	return &FilterByLanguageScope{Search: search}
}

func (s *FilterByLanguageScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Search == nil {
			return db
		} else {
			searchTerm := "%" + *s.Search + "%"
			return db.Where("books.language ILIKE ?", searchTerm)
		}
	}
}
