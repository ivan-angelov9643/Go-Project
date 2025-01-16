package managers

import (
	"awesomeProject/library-app/global/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=RatingManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type RatingManagerInterface interface {
	GetAll(*db.DBScope) ([]models.Rating, error)
	Get(uuid uuid.UUID) (models.Rating, error)
	Create(models.Rating) (models.Rating, error)
	Update(models.Rating) (models.Rating, error)
	Delete(uuid uuid.UUID) (models.Rating, error)
}
