package implementations

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReservationManager struct {
	db *gorm.DB
}

func NewReservationManager(db *gorm.DB) *ReservationManager {
	log.Info("[NewReservationManager] Initializing ReservationManager")

	return &ReservationManager{db}
}

func (m *ReservationManager) GetAll() ([]models.Reservation, *db_error.DBError) {
	log.Info("[ReservationManager.GetAll] Fetching all reservations")

	var allReservations []models.Reservation
	err := m.db.Find(&allReservations).Error
	if err != nil {
		log.Errorf("[ReservationManager.GetAll] Error fetching all reservations: %v", err)
		return nil, db_error.NewDBError(db_error.InternalError, "[ReservationManager.GetAll] Error fetching all reservations: %v", err)
	}

	log.Infof("[ReservationManager.GetAll] Successfully fetched all reservations")
	return allReservations, nil
}

func (m *ReservationManager) Get(idToGet uuid.UUID) (models.Reservation, *db_error.DBError) {
	log.Infof("[ReservationManager.Get] Fetching reservation with ID: %s", idToGet)

	var reservation models.Reservation
	err := m.db.First(&reservation, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[ReservationManager.Get] Error fetching reservation with ID %s: %v", idToGet, err)
		return models.Reservation{}, db_error.NewDBError(db_error.InternalError, "[ReservationManager.Get] Error fetching reservation with ID %s: %v", idToGet, err)
	}

	log.Infof("[ReservationManager.Get] Successfully fetched reservation with ID: %s", idToGet)
	return reservation, nil
}

func (m *ReservationManager) Create(newReservation models.Reservation) (models.Reservation, *db_error.DBError) {
	log.Infof("[ReservationManager.Create] Creating new reservation")

	err := newReservation.Validate()
	if err != nil {
		return models.Reservation{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	newReservation.ID = uuid.New()

	err = m.db.Create(&newReservation).Error
	if err != nil {
		log.Errorf("[ReservationManager.Create] Error creating new reservation with ID %s: %v", newReservation.ID, err)
		return models.Reservation{}, db_error.NewDBError(db_error.InternalError, "[ReservationManager.Create] Error creating new reservation with ID %s: %v", newReservation.ID, err)
	}

	log.Infof("[ReservationManager.Create] Successfully created reservation with ID: %s", newReservation.ID)
	return newReservation, nil
}

func (m *ReservationManager) Update(updatedReservation models.Reservation) (models.Reservation, *db_error.DBError) {
	log.Infof("[ReservationManager.Update] Updating reservation with ID: %s", updatedReservation.ID)

	err := updatedReservation.Validate()
	if err != nil {
		return models.Reservation{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	var reservation models.Reservation
	err = m.db.First(&reservation, "id = ?", updatedReservation.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[ReservationManager.Update] Reservation with ID %s does not exist", updatedReservation.ID)
			return models.Reservation{}, db_error.NewDBError(db_error.NotFoundError, "[ReservationManager.Update] Reservation with ID %s does not exist", updatedReservation.ID)
		}
		log.Errorf("[ReservationManager.Update] Error fetching reservation with ID %s: %v", updatedReservation.ID, err)
		return models.Reservation{}, db_error.NewDBError(db_error.InternalError, "[ReservationManager.Update] Error fetching reservation with ID %s: %v", updatedReservation.ID, err)
	}

	err = m.db.Model(&reservation).Updates(updatedReservation).Error
	if err != nil {
		log.Errorf("[ReservationManager.Update] Error updating reservation with ID %s: %v", updatedReservation.ID, err)
		return models.Reservation{}, db_error.NewDBError(db_error.InternalError, "[ReservationManager.Update] Error updating reservation with ID %s: %v", updatedReservation.ID, err)
	}

	log.Infof("[ReservationManager.Update] Successfully updated reservation with ID: %s", updatedReservation.ID)
	return updatedReservation, nil
}

func (m *ReservationManager) Delete(idToDelete uuid.UUID) (models.Reservation, *db_error.DBError) {
	log.Infof("[ReservationManager.Delete] Deleting reservation with ID: %s", idToDelete)

	var reservation models.Reservation
	err := m.db.First(&reservation, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[ReservationManager.Delete] Reservation with ID %s does not exist", idToDelete)
			return models.Reservation{}, db_error.NewDBError(db_error.NotFoundError, "[ReservationManager.Delete] Reservation with ID %s does not exist", idToDelete)
		}
		log.Errorf("[ReservationManager.Delete] Error fetching reservation with ID %s: %v", idToDelete, err)
		return models.Reservation{}, db_error.NewDBError(db_error.InternalError, "[ReservationManager.Delete] Error fetching reservation with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&reservation).Error
	if err != nil {
		log.Errorf("[ReservationManager.Delete] Error deleting reservation with ID %s: %v", idToDelete, err)
		return models.Reservation{}, db_error.NewDBError(db_error.InternalError, "[ReservationManager.Delete] Error deleting reservation with ID %s: %v", idToDelete, err)
	}

	log.Infof("[ReservationManager.Delete] Successfully deleted reservation with ID: %s", idToDelete)
	return reservation, nil
}
