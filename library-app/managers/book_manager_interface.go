package managers

import (
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
)

//go:generate mockery --name=BookManagerInterface --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type BookManagerInterface interface {
	GetAll(...db.DBScope) ([]models.Book, error)
	Get(uuid.UUID) (models.Book, error)
	Create(models.Book) (models.Book, error)
	Update(models.Book) (models.Book, error)
	Delete(uuid.UUID) (models.Book, error)
	Count(...db.DBScope) (int64, error)
}
