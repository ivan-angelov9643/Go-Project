package models

import (
	"awesomeProject/library-app/validation"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Book struct {
	BaseModel
	Title           string    `json:"title"`
	Year            int       `json:"year"`
	AuthorID        uuid.UUID `json:"author_id"`
	AuthorName      string    `json:"author_name" gorm:"<-:false"`
	CategoryID      uuid.UUID `json:"category_id"`
	CategoryName    string    `json:"category_name" gorm:"<-:false"`
	TotalCopies     int       `json:"total_copies"`
	AvailableCopies int       `json:"available_copies" gorm:"<-:false"`
	Language        string    `json:"language"`
}

func (book *Book) Validate(db *gorm.DB) error {
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

	var reservedCount, loanedCount int64

	err := db.Model(&Reservation{}).
		Where("book_id = ?", book.ID).
		Count(&reservedCount).Error
	if err != nil {
		log.Errorf("[Book.Validate] Failed to count reserved copies: %v", err)
		return err
	}

	err = db.Model(&Loan{}).
		Where("book_id = ? AND status = 'active'", book.ID).
		Count(&loanedCount).Error
	if err != nil {
		log.Errorf("[Book.Validate] Failed to count loaned copies: %v", err)
		return err
	}

	if book.TotalCopies < int(reservedCount+loanedCount) {
		err := fmt.Errorf("TotalCopies (%d) cannot be less than the sum of reserved (%d) and loaned (%d) copies",
			book.TotalCopies, reservedCount, loanedCount)
		log.Errorf("[Book.Validate] Validation failed: %v", err)
		return err
	}

	log.Info("[Book.Validate] Validation completed successfully")
	return nil
}
