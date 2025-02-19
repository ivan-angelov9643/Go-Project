package managers

import (
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
)

//go:generate mockery --name=ReservationManagerInterface --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ReservationManagerInterface interface {
	GetAll(...db.DBScope) ([]models.Reservation, error)
	Get(uuid.UUID) (models.Reservation, error)
	Create(models.Reservation) (models.Reservation, error)
	Update(models.Reservation) (models.Reservation, error)
	Delete(uuid.UUID) (models.Reservation, error)
	Count(...db.DBScope) (int64, error)
}
