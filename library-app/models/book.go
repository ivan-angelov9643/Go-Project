package models

import (
	"awesomeProject/library-app/validation"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Book struct {
	BaseModel
	Title           string    `json:"title"`
	Year            int       `json:"year"`
	AuthorID        uuid.UUID `json:"author_id"`
	CategoryID      uuid.UUID `json:"category_id"`
	TotalCopies     int       `json:"total_copies"`
	AvailableCopies int       `json:"available_copies" gorm:"-"`
	Language        string    `json:"language"`
}

func (book *Book) Validate() error {
	log.Info("[Book.Validate] Validating book data")

	for _, validationData := range validation.BookValidation {
		var fieldValue interface{}

		switch validationData.GetFieldName() {
		case "Title":
			fieldValue = book.Title
		case "Language":
			fieldValue = book.Language
		case "Year":
			fieldValue = book.Year
		case "Total Copies":
			fieldValue = book.TotalCopies
		}

		err := validationData.Validate(fieldValue)
		if err != nil {
			log.Errorf("[Book.Validate] Validation failed for field '%s' with value '%v': %v", validationData.GetFieldName(), fieldValue, err)
			return err
		}
	}

	log.Info("[Book.Validate] Validation completed successfully")
	return nil
}
