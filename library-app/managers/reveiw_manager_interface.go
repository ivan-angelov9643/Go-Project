package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=ReviewManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ReviewManagerInterface interface {
	GetAll() ([]models.Review, *db_error.DBError)
	Get(uuid uuid.UUID) (models.Review, *db_error.DBError)
	Create(models.Review) (models.Review, *db_error.DBError)
	Update(models.Review) (models.Review, *db_error.DBError)
	Delete(uuid uuid.UUID) (models.Review, *db_error.DBError)
}
