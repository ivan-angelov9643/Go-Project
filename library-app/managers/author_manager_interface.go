package managers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=AuthorManagerInterface --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type AuthorManagerInterface interface {
	GetAll(...db.DBScope) ([]models.Author, error)
	Get(uuid.UUID) (models.Author, error)
	Create(models.Author) (models.Author, error)
	Update(models.Author) (models.Author, error)
	Delete(uuid.UUID) (models.Author, error)
	Count(...db.DBScope) (int64, error)
}
