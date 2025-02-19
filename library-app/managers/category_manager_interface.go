package managers

import (
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
)

//go:generate mockery --name=CategoryManagerInterface --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type CategoryManagerInterface interface {
	GetAll(...db.DBScope) ([]models.Category, error)
	Get(uuid.UUID) (models.Category, error)
	Create(models.Category) (models.Category, error)
	Update(models.Category) (models.Category, error)
	Delete(uuid.UUID) (models.Category, error)
	Count(...db.DBScope) (int64, error)
}
