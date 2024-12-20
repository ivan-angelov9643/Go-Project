package managers

import (
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=CategoryManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type CategoryManagerInterface interface {
	GetAll() ([]models.Category, error)
	Get(uuid uuid.UUID) (models.Category, error)
	Create(models.Category) (models.Category, error)
	Update(models.Category) (models.Category, error)
	Delete(uuid uuid.UUID) (models.Category, error)
}
