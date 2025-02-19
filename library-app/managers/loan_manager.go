package managers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LoanManager struct {
	db *gorm.DB
}

func NewLoanManager(db *gorm.DB) *LoanManager {
	log.Info("[NewLoanManager] Initializing LoanManager")

	return &LoanManager{db}
}

func (m *LoanManager) GetAll(scopes ...db.DBScope) ([]models.Loan, error) {
	log.Info("[LoanManager.GetAll] Fetching all loans")

	var allLoans []models.Loan
	err := db.ApplyScopes(m.db, scopes).Table("loans").
		Select("loans.*, preferred_username as user_name, books.title as book_title").
		Joins("JOIN users ON users.id = loans.user_id").
		Joins("JOIN books ON books.id = loans.book_id").
		Find(&allLoans).Error
	if err != nil {
		log.Errorf("[LoanManager.GetAll] Error fetching all loans: %v", err)
		return nil, db.NewDBError(db.InternalError, "[LoanManager.GetAll] Error fetching all loans: %v", err)
	}

	log.Infof("[LoanManager.GetAll] Successfully fetched all loans")
	return allLoans, nil
}

func (m *LoanManager) Get(idToGet uuid.UUID) (models.Loan, error) {
	log.Infof("[LoanManager.Get] Fetching loan with ID: %s", idToGet)

	var loan models.Loan
	err := m.db.Table("loans").
		Select("loans.*, preferred_username as user_name, books.title as book_title").
		Joins("JOIN users ON users.id = loans.user_id").
		Joins("JOIN books ON books.id = loans.book_id").
		Where("loans.id = ?", idToGet).
		First(&loan).Error
	if err != nil {
		log.Errorf("[LoanManager.Get] Error fetching loan with ID %s: %v", idToGet, err)
		return models.Loan{}, db.NewDBError(db.InternalError, "[LoanManager.Get] Error fetching loan with ID %s: %v", idToGet, err)
	}

	log.Infof("[LoanManager.Get] Successfully fetched loan with ID: %s", idToGet)
	return loan, nil
}

func (m *LoanManager) Create(newLoan models.Loan) (models.Loan, error) {
	log.Infof("[LoanManager.Create] Creating new loan")

	err := newLoan.Validate()
	if err != nil {
		return models.Loan{}, db.NewDBError(db.ValidationError, err.Error())
	}

	newLoan.ID = uuid.New()

	err = m.db.Create(&newLoan).Error
	if err != nil {
		log.Errorf("[LoanManager.Create] Error creating new loan with ID %s: %v", newLoan.ID, err)
		return models.Loan{}, db.NewDBError(db.InternalError, "[LoanManager.Create] Error creating new loan with ID %s: %v", newLoan.ID, err)
	}

	log.Infof("[LoanManager.Create] Successfully created loan with ID: %s", newLoan.ID)
	return newLoan, nil
}

func (m *LoanManager) Update(updatedLoan models.Loan) (models.Loan, error) {
	log.Infof("[LoanManager.Update] Updating loan with ID: %s", updatedLoan.ID)

	err := updatedLoan.Validate()
	if err != nil {
		return models.Loan{}, db.NewDBError(db.ValidationError, err.Error())
	}

	var loan models.Loan
	err = m.db.First(&loan, "id = ?", updatedLoan.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[LoanManager.Update] Loan with ID %s does not exist", updatedLoan.ID)
			return models.Loan{}, db.NewDBError(db.NotFoundError, "[LoanManager.Update] Loan with ID %s does not exist", updatedLoan.ID)
		}
		log.Errorf("[LoanManager.Update] Error fetching loan with ID %s: %v", updatedLoan.ID, err)
		return models.Loan{}, db.NewDBError(db.InternalError, "[LoanManager.Update] Error fetching loan with ID %s: %v", updatedLoan.ID, err)
	}

	err = m.db.Model(&loan).Updates(updatedLoan).Error
	if err != nil {
		log.Errorf("[LoanManager.Update] Error updating loan with ID %s: %v", updatedLoan.ID, err)
		return models.Loan{}, db.NewDBError(db.InternalError, "[LoanManager.Update] Error updating loan with ID %s: %v", updatedLoan.ID, err)
	}

	log.Infof("[LoanManager.Update] Successfully updated loan with ID: %s", updatedLoan.ID)
	return updatedLoan, nil
}

func (m *LoanManager) Delete(idToDelete uuid.UUID) (models.Loan, error) {
	log.Infof("[LoanManager.Delete] Deleting loan with ID: %s", idToDelete)

	var loan models.Loan
	err := m.db.First(&loan, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[LoanManager.Delete] Loan with ID %s does not exist", idToDelete)
			return models.Loan{}, db.NewDBError(db.NotFoundError, "[LoanManager.Delete] Loan with ID %s does not exist", idToDelete)
		}
		log.Errorf("[LoanManager.Delete] Error fetching loan with ID %s: %v", idToDelete, err)
		return models.Loan{}, db.NewDBError(db.InternalError, "[LoanManager.Delete] Error fetching loan with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&loan).Error
	if err != nil {
		log.Errorf("[LoanManager.Delete] Error deleting loan with ID %s: %v", idToDelete, err)
		return models.Loan{}, db.NewDBError(db.InternalError, "[LoanManager.Delete] Error deleting loan with ID %s: %v", idToDelete, err)
	}

	log.Infof("[LoanManager.Delete] Successfully deleted loan with ID: %s", idToDelete)
	return loan, nil
}

func (m *LoanManager) Count(scopes ...db.DBScope) (int64, error) {
	log.Infof("[LoanManager.Count] Counting loans in the database")

	var count int64
	err := db.ApplyScopes(m.db, scopes).Table("loans").
		Select("loans.*, preferred_username as user_name, books.title as book_title").
		Joins("JOIN users ON users.id = loans.user_id").
		Joins("JOIN books ON books.id = loans.book_id").
		Count(&count).Error
	if err != nil {
		log.Errorf("[LoanManager.Count] Error counting loans: %v", err)
		return 0, db.NewDBError(db.InternalError, "[LoanManager.Count] Error counting loans: %v", err)
	}

	log.Infof("[LoanManager.Count] Successfully counted loans: %d", count)
	return count, nil
}
