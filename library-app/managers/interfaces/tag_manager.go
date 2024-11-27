package interfaces

import (
	"awesomeProject/library-app/models"
)

//go:generate mockery --name=TagManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type TagManager interface {
	GetAll() []models.Tag
	Get(string) (models.Tag, error)
	Create(models.Tag) (models.Tag, error)
	Delete(string) (models.Tag, error)
}
