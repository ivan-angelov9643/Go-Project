package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=AuthorManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type AuthorManagerInterface interface {
	GetAll() ([]models.Author, *db_error.DBError)
	Get(uuid uuid.UUID) (models.Author, *db_error.DBError)
	Create(models.Author) (models.Author, *db_error.DBError)
	Update(models.Author) (models.Author, *db_error.DBError)
	Delete(uuid uuid.UUID) (models.Author, *db_error.DBError)
}
