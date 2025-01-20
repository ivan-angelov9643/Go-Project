package managers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=LoanManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type LoanManagerInterface interface {
	GetAll(*db.AccessScope, *db.PagingScope) ([]models.Loan, error)
	Get(uuid uuid.UUID) (models.Loan, error)
	Create(models.Loan) (models.Loan, error)
	Update(models.Loan) (models.Loan, error)
	Delete(uuid uuid.UUID) (models.Loan, error)
}
