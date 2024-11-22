package interfaces

import (
	"awesomeProject/todo-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=ReviewManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ReviewManager interface {
	GetAll() []models.Review
	Get(uuid uuid.UUID) (models.Review, error)
	Create(models.Review) (models.Review, error)
	Update(models.Review) (models.Review, error)
	Delete(uuid uuid.UUID) (models.Review, error)
}
