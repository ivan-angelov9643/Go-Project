package interfaces

import (
	"awesomeProject/todo-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=ReservationManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ReservationManager interface {
	GetAll() []models.Reservation
	Get(uuid uuid.UUID) (models.Reservation, error)
	Create(models.Reservation) (models.Reservation, error)
	Update(models.Reservation) (models.Reservation, error)
	Delete(uuid uuid.UUID) (models.Reservation, error)
}
