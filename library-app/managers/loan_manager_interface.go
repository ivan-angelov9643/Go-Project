package managers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=LoanManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type LoanManagerInterface interface {
	GetAll(...db.DBScope) ([]models.Loan, error)
	Get(uuid.UUID) (models.Loan, error)
	Create(models.Loan) (models.Loan, error)
	Update(models.Loan) (models.Loan, error)
	Delete(uuid.UUID) (models.Loan, error)
	Count(...db.DBScope) (int64, error)
}
