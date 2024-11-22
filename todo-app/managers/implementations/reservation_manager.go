package implementations

import (
	"awesomeProject/todo-app/models"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReservationManager struct {
	db           *gorm.DB
	reservations map[uuid.UUID]models.Reservation
}

func NewReservationManager(db *gorm.DB) *ReservationManager {
	log.Info("[NewReservationManager] Initializing ReservationManager")

	return &ReservationManager{db, make(map[uuid.UUID]models.Reservation)}
}

func (m *ReservationManager) GetAll() []models.Reservation {
	log.Info("[ReservationManager.GetAll] Fetching all reservations")

	allReservations := make([]models.Reservation, 0, len(m.reservations))
	for _, reservation := range m.reservations {
		allReservations = append(allReservations, reservation)
	}

	log.Infof("[ReservationManager.GetAll] Successfully fetched all reservations")
	return allReservations
}

func (m *ReservationManager) Get(idToGet uuid.UUID) (models.Reservation, error) {
	log.Infof("[ReservationManager.Get] Fetching reservation with ID: %s", idToGet)

	reservation, exists := m.reservations[idToGet]
	if !exists {
		log.Errorf("[ReservationManager.Get] Reservation with ID %s not found", idToGet)
		return models.Reservation{}, fmt.Errorf("[ReservationManager.Get] Reservation with ID %s not found", idToGet)
	}

	log.Infof("[ReservationManager.Get] Successfully fetched reservation with ID: %s", idToGet)
	return reservation, nil
}

func (m *ReservationManager) Create(newReservation models.Reservation) (models.Reservation, error) {
	log.Infof("[ReservationManager.Create] Creating new reservation")

	if newReservation.ID == uuid.Nil {
		newReservation.ID = uuid.New()
	}

	_, exists := m.reservations[newReservation.ID]
	if exists {
		log.Errorf("[ReservationManager.Create] Reservation with ID %s already exists", newReservation.ID)
		return models.Reservation{}, fmt.Errorf("[ReservationManager.Create] Reservation with ID %s already exists", newReservation.ID)
	}

	m.reservations[newReservation.ID] = newReservation
	log.Infof("[ReservationManager.Create] Successfully created reservation with ID: %s", newReservation.ID)
	return newReservation, nil
}

func (m *ReservationManager) Update(updatedReservation models.Reservation) (models.Reservation, error) {
	log.Infof("[ReservationManager.Update] Updating reservation with ID: %s", updatedReservation.ID)

	_, exists := m.reservations[updatedReservation.ID]
	if !exists {
		log.Errorf("[ReservationManager.Update] Reservation with ID %s not found", updatedReservation.ID)
		return models.Reservation{}, fmt.Errorf("[ReservationManager.Update] Reservation with ID %s not found", updatedReservation.ID)
	}

	m.reservations[updatedReservation.ID] = updatedReservation
	log.Infof("[ReservationManager.Update] Successfully updated reservation with ID: %s", updatedReservation.ID)
	return updatedReservation, nil
}

func (m *ReservationManager) Delete(idToDelete uuid.UUID) (models.Reservation, error) {
	log.Infof("[ReservationManager.Delete] Deleting reservation with ID: %s", idToDelete)

	deletedReservation, exists := m.reservations[idToDelete]
	if !exists {
		log.Errorf("[ReservationManager.Delete] Reservation with ID %s not found", idToDelete)
		return models.Reservation{}, fmt.Errorf("[ReservationManager.Delete] Reservation with ID %s not found", idToDelete)
	}

	delete(m.reservations, idToDelete)
	log.Infof("[ReservationManager.Delete] Successfully deleted reservation with ID: %s", idToDelete)
	return deletedReservation, nil
}
