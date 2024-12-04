package managers

import (
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=ReviewManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ReviewManagerInterface interface {
	GetAll() ([]models.Review, error)
	Get(uuid uuid.UUID) (models.Review, error)
	Create(models.Review) (models.Review, error)
	Update(models.Review) (models.Review, error)
	Delete(uuid uuid.UUID) (models.Review, error)
}
