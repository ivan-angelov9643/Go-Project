package models

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type Loan struct {
	BaseModel
	UserID     uuid.UUID  `json:"user_id"`
	UserName   string     `json:"user_name" gorm:"<-:false"`
	BookID     uuid.UUID  `json:"book_id"`
	BookTitle  string     `json:"book_title" gorm:"<-:false"`
	StartDate  time.Time  `json:"start_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"` // nullable
	Status     string     `json:"status"`                // ENUM: active, completed
}

func (loan *Loan) Validate() error {
	log.Info("[Loan.Validate] Validating loan data")
	// Add validation logic here
	log.Info("[Loan.Validate] Validation completed successfully")
	return nil
}
