package db

import (
	"gorm.io/gorm"
	"net/http"
)

type FilterByStatusScope struct {
	Status *string
}

func NewFilterByStatusScope(r *http.Request) *FilterByStatusScope {
	statusParam := r.URL.Query().Get("status")

	var status *string = nil
	if statusParam != "" && statusParam != "null" {
		status = &statusParam
	}

	return &FilterByStatusScope{Status: status}
}

func (s *FilterByStatusScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Status == nil || *s.Status == "all" {
			return db
		} else {
			return db.Where("status = ?", s.Status)
		}
	}
}
