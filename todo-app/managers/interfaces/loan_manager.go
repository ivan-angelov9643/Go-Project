package interfaces

import (
	"awesomeProject/todo-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=LoanManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type LoanManager interface {
	GetAll() []models.Loan
	Get(uuid uuid.UUID) (models.Loan, error)
	Create(models.Loan) (models.Loan, error)
	Update(models.Loan) (models.Loan, error)
	Delete(uuid uuid.UUID) (models.Loan, error)
}
