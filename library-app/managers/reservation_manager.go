package managers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/global"
	"github.com/ivan-angelov9643/go-project/library-app/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type ReservationManager struct {
	db *gorm.DB
}

func NewReservationManager(db *gorm.DB) *ReservationManager {
	log.Info("[NewReservationManager] Initializing ReservationManager")

	return &ReservationManager{db}
}

func (m *ReservationManager) CleanupExpiredReservations() {
	log.Info("[ReservationManager.CleanupExpiredReservations] Starting cleanup job")
	err := m.db.Where("expiry_date <= ?", time.Now()).Delete(&models.Reservation{}).Error
	if err != nil {
		log.Errorf("[ReservationManager.CleanupExpiredReservations] Error cleaning up expired reservations: %v", err)
	} else {
		log.Info("[ReservationManager.CleanupExpiredReservations] Cleanup job completed")
	}
}

func (m *ReservationManager) GetAll(scopes ...db.DBScope) ([]models.Reservation, error) {
	log.Info("[ReservationManager.GetAll] Fetching all reservations")

	var allReservations []models.Reservation
	err := db.ApplyScopes(m.db, scopes).Table("reservations").
		Select("reservations.*, preferred_username as user_name, books.title as book_title").
		Joins("JOIN users ON users.id = reservations.user_id").
		Joins("JOIN books ON books.id = reservations.book_id").
		Find(&allReservations).Error
	if err != nil {
		log.Errorf("[ReservationManager.GetAll] Error fetching all reservations: %v", err)
		return nil, db.NewDBError(db.InternalError, "[ReservationManager.GetAll] Error fetching all reservations: %v", err)
	}

	log.Infof("[ReservationManager.GetAll] Successfully fetched all reservations")
	return allReservations, nil
}

func (m *ReservationManager) Get(idToGet uuid.UUID) (models.Reservation, error) {
	log.Infof("[ReservationManager.Get] Fetching reservation with ID: %s", idToGet)

	var reservation models.Reservation
	err := m.db.Table("reservations").
		Select("reservations.*, preferred_username as user_name, books.title as book_title").
		Joins("JOIN users ON users.id = reservations.user_id").
		Joins("JOIN books ON books.id = reservations.book_id").
		Where("reservations.id = ?", idToGet).
		First(&reservation).Error
	if err != nil {
		log.Errorf("[ReservationManager.Get] Error fetching reservation with ID %s: %v", idToGet, err)
		return models.Reservation{}, db.NewDBError(db.InternalError, "[ReservationManager.Get] Error fetching reservation with ID %s: %v", idToGet, err)
	}

	log.Infof("[ReservationManager.Get] Successfully fetched reservation with ID: %s", idToGet)
	return reservation, nil
}

func (m *ReservationManager) Create(newReservation models.Reservation) (models.Reservation, error) {
	log.Infof("[ReservationManager.Create] Creating new reservation")

	err := newReservation.Validate()
	if err != nil {
		return models.Reservation{}, db.NewDBError(db.ValidationError, err.Error())
	}

	newReservation.ID = uuid.New()
	newReservation.ExpiryDate = time.Now().Add(global.ReservationDuration)

	err = m.db.Create(&newReservation).Error
	if err != nil {
		log.Errorf("[ReservationManager.Create] Error creating new reservation with ID %s: %v", newReservation.ID, err)
		return models.Reservation{}, db.NewDBError(db.InternalError, "[ReservationManager.Create] Error creating new reservation with ID %s: %v", newReservation.ID, err)
	}

	log.Infof("[ReservationManager.Create] Successfully created reservation with ID: %s", newReservation.ID)
	return newReservation, nil
}

func (m *ReservationManager) Update(updatedReservation models.Reservation) (models.Reservation, error) {
	log.Infof("[ReservationManager.Update] Updating reservation with ID: %s", updatedReservation.ID)

	err := updatedReservation.Validate()
	if err != nil {
		return models.Reservation{}, db.NewDBError(db.ValidationError, err.Error())
	}

	var reservation models.Reservation
	err = m.db.First(&reservation, "id = ?", updatedReservation.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[ReservationManager.Update] Reservation with ID %s does not exist", updatedReservation.ID)
			return models.Reservation{}, db.NewDBError(db.NotFoundError, "[ReservationManager.Update] Reservation with ID %s does not exist", updatedReservation.ID)
		}
		log.Errorf("[ReservationManager.Update] Error fetching reservation with ID %s: %v", updatedReservation.ID, err)
		return models.Reservation{}, db.NewDBError(db.InternalError, "[ReservationManager.Update] Error fetching reservation with ID %s: %v", updatedReservation.ID, err)
	}

	err = m.db.Model(&reservation).Updates(updatedReservation).Error
	if err != nil {
		log.Errorf("[ReservationManager.Update] Error updating reservation with ID %s: %v", updatedReservation.ID, err)
		return models.Reservation{}, db.NewDBError(db.InternalError, "[ReservationManager.Update] Error updating reservation with ID %s: %v", updatedReservation.ID, err)
	}

	log.Infof("[ReservationManager.Update] Successfully updated reservation with ID: %s", updatedReservation.ID)
	return updatedReservation, nil
}

func (m *ReservationManager) Delete(idToDelete uuid.UUID) (models.Reservation, error) {
	log.Infof("[ReservationManager.Delete] Deleting reservation with ID: %s", idToDelete)

	var reservation models.Reservation
	err := m.db.First(&reservation, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[ReservationManager.Delete] Reservation with ID %s does not exist", idToDelete)
			return models.Reservation{}, db.NewDBError(db.NotFoundError, "[ReservationManager.Delete] Reservation with ID %s does not exist", idToDelete)
		}
		log.Errorf("[ReservationManager.Delete] Error fetching reservation with ID %s: %v", idToDelete, err)
		return models.Reservation{}, db.NewDBError(db.InternalError, "[ReservationManager.Delete] Error fetching reservation with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&reservation).Error
	if err != nil {
		log.Errorf("[ReservationManager.Delete] Error deleting reservation with ID %s: %v", idToDelete, err)
		return models.Reservation{}, db.NewDBError(db.InternalError, "[ReservationManager.Delete] Error deleting reservation with ID %s: %v", idToDelete, err)
	}

	log.Infof("[ReservationManager.Delete] Successfully deleted reservation with ID: %s", idToDelete)
	return reservation, nil
}

func (m *ReservationManager) Count(scopes ...db.DBScope) (int64, error) {
	log.Infof("[ReservationManager.Count] Counting reservations in the database")

	var count int64
	err := db.ApplyScopes(m.db, scopes).Table("reservations").
		Select("reservations.*, preferred_username as user_name, books.title as book_title").
		Joins("JOIN users ON users.id = reservations.user_id").
		Joins("JOIN books ON books.id = reservations.book_id").
		Count(&count).Error
	if err != nil {
		log.Errorf("[ReservationManager.Count] Error counting reservations: %v", err)
		return 0, db.NewDBError(db.InternalError, "[ReservationManager.Count] Error counting reservations: %v", err)
	}

	log.Infof("[ReservationManager.Count] Successfully counted reservations: %d", count)
	return count, nil
}
