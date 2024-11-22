package models

import (
	"github.com/google/uuid"
	"time"
)

type Reservation struct {
	BaseModel
	UserID     uuid.UUID `json:"user_id"`
	BookID     uuid.UUID `json:"book_id"`
	ExpiryDate time.Time `json:"expiry_date"`
}
