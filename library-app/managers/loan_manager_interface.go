package managers

import (
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
)

//go:generate mockery --name=LoanManagerInterface --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type LoanManagerInterface interface {
	GetAll(...db.DBScope) ([]models.Loan, error)
	Get(uuid.UUID) (models.Loan, error)
	Create(models.Loan) (models.Loan, error)
	Update(models.Loan) (models.Loan, error)
	Delete(uuid.UUID) (models.Loan, error)
	Count(...db.DBScope) (int64, error)
}
