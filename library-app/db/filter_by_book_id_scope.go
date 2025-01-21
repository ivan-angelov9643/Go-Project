package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type FilterByBookIDScope struct {
	BookID *uuid.UUID
}

func NewRatingFilterScope(r *http.Request) *FilterByBookIDScope {
	bookIDParam := r.URL.Query().Get("book_id")

	var bookID *uuid.UUID = nil
	if bookIDParam != "" && bookIDParam != "null" {
		parsedBookID, err := uuid.Parse(bookIDParam)
		if err == nil {
			bookID = &parsedBookID
		}
	}
	return &FilterByBookIDScope{bookID}
}

func (s *FilterByBookIDScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.BookID == nil {
			return db
		} else {
			return db.Where("book_id = ?", s.BookID)
		}
	}
}
