package models

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type Reservation struct {
	BaseModel
	UserID     uuid.UUID `json:"user_id"`
	UserName   string    `json:"user_name" gorm:"<-:false"`
	BookID     uuid.UUID `json:"book_id"`
	BookTitle  string    `json:"book_title" gorm:"<-:false"`
	ExpiryDate time.Time `json:"expiry_date"`
}

func (reservation *Reservation) Validate() error {
	log.Info("[Reservation.Validate] Validating reservation data")

	if reservation.CreatedAt.After(reservation.ExpiryDate) {
		err := fmt.Errorf("CreatedAt cannot be after ExpiryDate: CreatedAt = %v, ExpiryDate = %v", reservation.CreatedAt, reservation.ExpiryDate)
		log.Errorf("[Reservation.Validate] Validation failed: %v", err)
		return err
	}

	log.Info("[Reservation.Validate] Validation completed successfully")
	return nil
}
