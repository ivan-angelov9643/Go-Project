package managers

import (
	"awesomeProject/library-app/global/db"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=ReservationManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ReservationManagerInterface interface {
	GetAll(*db.DBScope) ([]models.Reservation, error)
	Get(uuid uuid.UUID) (models.Reservation, error)
	Create(models.Reservation) (models.Reservation, error)
	Update(models.Reservation) (models.Reservation, error)
	Delete(uuid uuid.UUID) (models.Reservation, error)
}
