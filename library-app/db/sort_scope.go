package db

import (
	"gorm.io/gorm"
	"net/http"
)

type SortScope struct {
	SortBy    *string
	SortOrder *string
}

func NewSortScope(r *http.Request) *SortScope {
	sortByParam := r.URL.Query().Get("sort_by")
	sortOrderParam := r.URL.Query().Get("sort_order")

	var sortBy *string = nil
	var sortOrder *string = nil

	if sortByParam != "" && sortByParam != "null" {
		sortBy = &sortByParam
	}

	if sortOrderParam != "" && sortOrderParam != "null" {
		sortOrder = &sortOrderParam
	}

	return &SortScope{SortBy: sortBy, SortOrder: sortOrder}
}

func (s *SortScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.SortBy == nil || s.SortOrder == nil {
			return db
		} else {
			if *s.SortOrder == "desc" {
				return db.Order(*s.SortBy + " desc")
			}
			return db.Order(*s.SortBy + " asc")
		}
	}
}
