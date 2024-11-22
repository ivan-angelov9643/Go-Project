package models

import (
	"github.com/google/uuid"
	"time"
)

type Loan struct {
	BaseModel
	UserID     uuid.UUID  `json:"user_id"`
	BookID     uuid.UUID  `json:"book_id"`
	StartDate  time.Time  `json:"start_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"` // nullable
	Status     string     `json:"status"`                // ENUM: active, completed
}
