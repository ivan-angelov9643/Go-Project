package implementations

import (
	"awesomeProject/todo-app/models"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LoanManager struct {
	db    *gorm.DB
	loans map[uuid.UUID]models.Loan
}

func NewLoanManager(db *gorm.DB) *LoanManager {
	log.Info("[NewLoanManager] Initializing LoanManager")

	return &LoanManager{db, make(map[uuid.UUID]models.Loan)}
}

func (m *LoanManager) GetAll() []models.Loan {
	log.Info("[LoanManager.GetAll] Fetching all loans")

	allLoans := make([]models.Loan, 0, len(m.loans))
	for _, loan := range m.loans {
		allLoans = append(allLoans, loan)
	}

	log.Infof("[LoanManager.GetAll] Successfully fetched all loans")
	return allLoans
}

func (m *LoanManager) Get(idToGet uuid.UUID) (models.Loan, error) {
	log.Infof("[LoanManager.Get] Fetching loan with ID: %s", idToGet)

	loan, exists := m.loans[idToGet]
	if !exists {
		log.Errorf("[LoanManager.Get] Loan with ID %s not found", idToGet)
		return models.Loan{}, fmt.Errorf("[LoanManager.Get] Loan with ID %s not found", idToGet)
	}

	log.Infof("[LoanManager.Get] Successfully fetched loan with ID: %s", idToGet)
	return loan, nil
}

func (m *LoanManager) Create(newLoan models.Loan) (models.Loan, error) {
	log.Infof("[LoanManager.Create] Creating new loan")

	if newLoan.ID == uuid.Nil {
		newLoan.ID = uuid.New()
	}

	_, exists := m.loans[newLoan.ID]
	if exists {
		log.Errorf("[LoanManager.Create] Loan with ID %s already exists", newLoan.ID)
		return models.Loan{}, fmt.Errorf("[LoanManager.Create] Loan with ID %s already exists", newLoan.ID)
	}

	m.loans[newLoan.ID] = newLoan
	log.Infof("[LoanManager.Create] Successfully created loan with ID: %s", newLoan.ID)
	return newLoan, nil
}

func (m *LoanManager) Update(updatedLoan models.Loan) (models.Loan, error) {
	log.Infof("[LoanManager.Update] Updating loan with ID: %s", updatedLoan.ID)

	_, exists := m.loans[updatedLoan.ID]
	if !exists {
		log.Errorf("[LoanManager.Update] Loan with ID %s not found", updatedLoan.ID)
		return models.Loan{}, fmt.Errorf("[LoanManager.Update] Loan with ID %s not found", updatedLoan.ID)
	}

	m.loans[updatedLoan.ID] = updatedLoan
	log.Infof("[LoanManager.Update] Successfully updated loan with ID: %s", updatedLoan.ID)
	return updatedLoan, nil
}

func (m *LoanManager) Delete(idToDelete uuid.UUID) (models.Loan, error) {
	log.Infof("[LoanManager.Delete] Deleting loan with ID: %s", idToDelete)

	deletedLoan, exists := m.loans[idToDelete]
	if !exists {
		log.Errorf("[LoanManager.Delete] Loan with ID %s not found", idToDelete)
		return models.Loan{}, fmt.Errorf("[LoanManager.Delete] Loan with ID %s not found", idToDelete)
	}

	delete(m.loans, idToDelete)
	log.Infof("[LoanManager.Delete] Successfully deleted loan with ID: %s", idToDelete)
	return deletedLoan, nil
}
