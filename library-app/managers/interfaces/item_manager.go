package interfaces

import (
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=ItemManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ItemManager interface {
	GetAll() []models.Item
	Get(uuid.UUID) (models.Item, error)
	Create(models.Item) (models.Item, error)
	Update(models.Item) (models.Item, error)
	Delete(uuid.UUID) (models.Item, error)
}
