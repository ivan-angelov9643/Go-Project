package managers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=BookManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type BookManagerInterface interface {
	GetAll(*db.AccessScope, *db.PagingScope) ([]models.Book, error)
	Get(uuid.UUID) (models.Book, error)
	Create(models.Book) (models.Book, error)
	Update(models.Book) (models.Book, error)
	Delete(uuid.UUID) (models.Book, error)
	Count(*db.AccessScope) (int64, error)
}
