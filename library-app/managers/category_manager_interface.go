package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=CategoryManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type CategoryManagerInterface interface {
	GetAll() ([]models.Category, *db_error.DBError)
	Get(uuid uuid.UUID) (models.Category, *db_error.DBError)
	Create(models.Category) (models.Category, *db_error.DBError)
	Update(models.Category) (models.Category, *db_error.DBError)
	Delete(uuid uuid.UUID) (models.Category, *db_error.DBError)
}
