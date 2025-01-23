package db

import (
	"gorm.io/gorm"
	"net/http"
)

type FilterByTitleScope struct {
	Search *string
}

func NewFilterByTitleScope(r *http.Request) *FilterByTitleScope {
	searchParam := r.URL.Query().Get("title")

	var search *string = nil
	if searchParam != "" && searchParam != "null" {
		search = &searchParam
	}

	return &FilterByTitleScope{Search: search}
}

func (s *FilterByTitleScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Search == nil {
			return db
		} else {
			searchTerm := "%" + *s.Search + "%"
			return db.Where("books.title ILIKE ?", searchTerm)
		}
	}
}
