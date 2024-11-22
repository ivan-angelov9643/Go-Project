package models

import (
	"github.com/google/uuid"
)

type Book struct {
	BaseModel
	Title       string    `json:"title"`
	Year        int       `json:"year"`
	AuthorID    uuid.UUID `json:"author_id"`
	CategoryID  uuid.UUID `json:"category_id"`
	TotalCopies int       `json:"total_copies"`
	Language    string    `json:"language"`
}
