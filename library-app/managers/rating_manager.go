package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RatingManager struct {
	db *gorm.DB
}

func NewRatingManager(db *gorm.DB) *RatingManager {
	log.Info("[NewRatingManager] Initializing RatingManager")

	return &RatingManager{db}
}

func (m *RatingManager) GetAll() ([]models.Rating, error) {
	log.Info("[RatingManager.GetAll] Fetching all ratings")

	var allRatings []models.Rating
	err := m.db.Find(&allRatings).Error
	if err != nil {
		log.Errorf("[RatingManager.GetAll] Error fetching all ratings: %v", err)
		return nil, db_error.NewDBError(db_error.InternalError, "[RatingManager.GetAll] Error fetching all ratings: %v", err)
	}

	log.Infof("[RatingManager.GetAll] Successfully fetched all ratings")
	return allRatings, nil
}

func (m *RatingManager) Get(idToGet uuid.UUID) (models.Rating, error) {
	log.Infof("[RatingManager.Get] Fetching rating with ID: %s", idToGet)

	var rating models.Rating
	err := m.db.First(&rating, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[RatingManager.Get] Error fetching rating with ID %s: %v", idToGet, err)
		return models.Rating{}, db_error.NewDBError(db_error.InternalError, "[RatingManager.Get] Error fetching rating with ID %s: %v", idToGet, err)
	}

	log.Infof("[RatingManager.Get] Successfully fetched rating with ID: %s", idToGet)
	return rating, nil
}

func (m *RatingManager) Create(newRating models.Rating) (models.Rating, error) {
	log.Infof("[RatingManager.Create] Creating new rating")

	err := newRating.Validate()
	if err != nil {
		return models.Rating{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	newRating.ID = uuid.New()

	err = m.db.Create(&newRating).Error
	if err != nil {
		log.Errorf("[RatingManager.Create] Error creating new rating with ID %s: %v", newRating.ID, err)
		return models.Rating{}, db_error.NewDBError(db_error.InternalError, "[RatingManager.Create] Error creating new rating with ID %s: %v", newRating.ID, err)
	}

	log.Infof("[RatingManager.Create] Successfully created rating with ID: %s", newRating.ID)
	return newRating, nil
}

func (m *RatingManager) Update(updatedRating models.Rating) (models.Rating, error) {
	log.Infof("[RatingManager.Update] Updating rating with ID: %s", updatedRating.ID)

	err := updatedRating.Validate()
	if err != nil {
		return models.Rating{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	var rating models.Rating
	err = m.db.First(&rating, "id = ?", updatedRating.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[RatingManager.Update] Rating with ID %s does not exist", updatedRating.ID)
			return models.Rating{}, db_error.NewDBError(db_error.NotFoundError, "[RatingManager.Update] Rating with ID %s does not exist", updatedRating.ID)
		}
		log.Errorf("[RatingManager.Update] Error fetching rating with ID %s: %v", updatedRating.ID, err)
		return models.Rating{}, db_error.NewDBError(db_error.InternalError, "[RatingManager.Update] Error fetching rating with ID %s: %v", updatedRating.ID, err)
	}

	err = m.db.Model(&rating).Updates(updatedRating).Error
	if err != nil {
		log.Errorf("[RatingManager.Update] Error updating rating with ID %s: %v", updatedRating.ID, err)
		return models.Rating{}, db_error.NewDBError(db_error.InternalError, "[RatingManager.Update] Error updating rating with ID %s: %v", updatedRating.ID, err)
	}

	log.Infof("[RatingManager.Update] Successfully updated rating with ID: %s", updatedRating.ID)
	return updatedRating, nil
}

func (m *RatingManager) Delete(idToDelete uuid.UUID) (models.Rating, error) {
	log.Infof("[RatingManager.Delete] Deleting rating with ID: %s", idToDelete)

	var rating models.Rating
	err := m.db.First(&rating, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[RatingManager.Delete] Rating with ID %s does not exist", idToDelete)
			return models.Rating{}, db_error.NewDBError(db_error.NotFoundError, "[RatingManager.Delete] Rating with ID %s does not exist", idToDelete)
		}
		log.Errorf("[RatingManager.Delete] Error fetching rating with ID %s: %v", idToDelete, err)
		return models.Rating{}, db_error.NewDBError(db_error.InternalError, "[RatingManager.Delete] Error fetching rating with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&rating).Error
	if err != nil {
		log.Errorf("[RatingManager.Delete] Error deleting rating with ID %s: %v", idToDelete, err)
		return models.Rating{}, db_error.NewDBError(db_error.InternalError, "[RatingManager.Delete] Error deleting rating with ID %s: %v", idToDelete, err)
	}

	log.Infof("[RatingManager.Delete] Successfully deleted rating with ID: %s", idToDelete)
	return rating, nil
}
