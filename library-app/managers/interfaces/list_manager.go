package interfaces

import (
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=ListManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ListManager interface {
	GetAll() []models.List
	Get(uuid.UUID) (models.List, error)
	Create(models.List) (models.List, error)
	Update(models.List) (models.List, error)
	Delete(uuid.UUID) (models.List, error)
}
