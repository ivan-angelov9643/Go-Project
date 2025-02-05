package managers

import (
	"awesomeProject/library-app/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
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
