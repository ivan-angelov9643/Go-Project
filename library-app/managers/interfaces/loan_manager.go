package interfaces

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=LoanManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type LoanManager interface {
	GetAll() ([]models.Loan, *db_error.DBError)
	Get(uuid uuid.UUID) (models.Loan, *db_error.DBError)
	Create(models.Loan) (models.Loan, *db_error.DBError)
	Update(models.Loan) (models.Loan, *db_error.DBError)
	Delete(uuid uuid.UUID) (models.Loan, *db_error.DBError)
}
