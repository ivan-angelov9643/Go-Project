package db

import (
	"gorm.io/gorm"
	"net/http"
)

type FilterByCategoryNameScope struct {
	Search *string
}

func NewFilterByCategoryNameScope(r *http.Request) *FilterByCategoryNameScope {
	searchParam := r.URL.Query().Get("category_name")

	var search *string = nil
	if searchParam != "" && searchParam != "null" {
		search = &searchParam
	}

	return &FilterByCategoryNameScope{Search: search}
}

func (s *FilterByCategoryNameScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Search == nil {
			return db
		} else {
			searchTerm := "%" + *s.Search + "%"
			return db.Where("categories.name ILIKE ?", searchTerm)
		}
	}
}
