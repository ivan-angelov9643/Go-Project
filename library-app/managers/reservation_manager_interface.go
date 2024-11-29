package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=ReservationManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ReservationManagerInterface interface {
	GetAll() ([]models.Reservation, *db_error.DBError)
	Get(uuid uuid.UUID) (models.Reservation, *db_error.DBError)
	Create(models.Reservation) (models.Reservation, *db_error.DBError)
	Update(models.Reservation) (models.Reservation, *db_error.DBError)
	Delete(uuid uuid.UUID) (models.Reservation, *db_error.DBError)
}
