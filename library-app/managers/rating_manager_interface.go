package managers

import (
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
)

//go:generate mockery --name=RatingManagerInterface --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type RatingManagerInterface interface {
	GetAll(...db.DBScope) ([]models.Rating, error)
	Get(uuid.UUID) (models.Rating, error)
	Create(models.Rating) (models.Rating, error)
	Update(models.Rating) (models.Rating, error)
	Delete(uuid.UUID) (models.Rating, error)
	Count(...db.DBScope) (int64, error)
}
