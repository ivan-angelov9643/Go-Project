package models

import (
	"github.com/google/uuid"
)

type Review struct {
	BaseModel
	UserID  uuid.UUID `json:"user_id"`
	BookID  uuid.UUID `json:"book_id"`
	Content string    `json:"content"`
	Rating  float64   `json:"rating"` // 1 to 5
}
