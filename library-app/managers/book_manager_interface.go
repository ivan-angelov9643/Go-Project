package managers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=BookManagerInterface --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type BookManagerInterface interface {
	GetAll(...db.DBScope) ([]models.Book, error)
	Get(uuid.UUID) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(models.Book) (models.Book, error)
	Delete(uuid.UUID) (models.Book, error)
	Count(...db.DBScope) (int64, error)
}
