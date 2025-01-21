package managers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=UserManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type UserManagerInterface interface {
	GetAll(*db.AccessScope, *db.PagingScope) ([]models.User, error)
	Get(uuid.UUID) (models.User, error)
	Create(models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(uuid.UUID) (models.User, error)
	Count(*db.AccessScope) (int64, error)
}
