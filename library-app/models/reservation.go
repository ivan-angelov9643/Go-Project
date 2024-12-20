package models

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type Reservation struct {
	BaseModel
	UserID     uuid.UUID `json:"user_id"`
	BookID     uuid.UUID `json:"book_id"`
	ExpiryDate time.Time `json:"expiry_date"`
}

func (reservation *Reservation) Validate() error {
	log.Info("[Reservation.Validate] Validating reservation data")
	// Add validation logic here
	log.Info("[Reservation.Validate] Validation completed successfully")
	return nil
}
