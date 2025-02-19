package managers

import (
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
)

//go:generate mockery --name=UserManagerInterface --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type UserManagerInterface interface {
	GetAll(...db.DBScope) ([]models.User, error)
	Get(uuid.UUID) (models.User, error)
	Create(models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(uuid.UUID) (models.User, error)
	Count(...db.DBScope) (int64, error)
}
