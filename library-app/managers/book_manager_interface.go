package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=BookManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type BookManagerInterface interface {
	GetAll() ([]models.Book, *db_error.DBError)
	Get(uuid uuid.UUID) (models.Book, *db_error.DBError)
	Create(models.Book) (models.Book, *db_error.DBError)
	Update(models.Book) (models.Book, *db_error.DBError)
	Delete(uuid uuid.UUID) (models.Book, *db_error.DBError)
}
