package models

import (
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/validation"
	log "github.com/sirupsen/logrus"
)

type Rating struct {
	BaseModel
	UserID    uuid.UUID `json:"user_id"`
	UserName  string    `json:"user_name" gorm:"<-:false"`
	BookID    uuid.UUID `json:"book_id"`
	BookTitle string    `json:"book_title" gorm:"<-:false"`
	Content   string    `json:"content"`
	Value     int       `json:"value"` // 1 to 5
}

func (rating *Rating) Validate() error {
	log.Info("[Rating.Validate] Validating rating data")

	for _, validationData := range validation.RatingValidation {
		var fieldValue interface{}

		switch validationData.GetFieldName() {
		case "Content":
			fieldValue = rating.Content
		case "Value":
			fieldValue = rating.Value
		}

		err := validationData.Validate(fieldValue)
		if err != nil {
			log.Errorf("[Rating.Validate] Validation failed for field '%s' with value '%v': %v", validationData.GetFieldName(), fieldValue, err)
			return err
		}
	}

	log.Info("[Rating.Validate] Validation completed successfully")
	return nil
}
