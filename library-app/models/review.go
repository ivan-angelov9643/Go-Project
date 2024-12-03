package models

import (
	"awesomeProject/library-app/global/validation"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Review struct {
	BaseModel
	UserID  uuid.UUID `json:"user_id"`
	BookID  uuid.UUID `json:"book_id"`
	Content string    `json:"content"`
	Rating  float64   `json:"rating"` // 1 to 5
}

func (review *Review) Validate() error {
	log.Info("[Review.Validate] Validating review data")

	for _, validationData := range validation.ReviewValidation {
		var fieldValue interface{}

		switch validationData.GetFieldName() {
		case "Content":
			fieldValue = review.Content
		case "Rating":
			fieldValue = review.Rating
		}

		err := validationData.Validate(fieldValue)
		if err != nil {
			log.Errorf("[Review.Validate] Validation failed for field '%s' with value '%v': %v", validationData.GetFieldName(), fieldValue, err)
			return err
		}
	}

	log.Info("[Review.Validate] Validation completed successfully")
	return nil
}
