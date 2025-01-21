package managers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=AuthorManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type AuthorManagerInterface interface {
	GetAll(*db.AccessScope, *db.PagingScope) ([]models.Author, error)
	Get(uuid.UUID) (models.Author, error)
	Create(models.Author) (models.Author, error)
	Update(models.Author) (models.Author, error)
	Delete(uuid uuid.UUID) (models.Author, error)
	Count(*db.AccessScope) (int64, error)
}
