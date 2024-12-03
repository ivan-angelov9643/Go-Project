package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReviewManager struct {
	db *gorm.DB
}

func NewReviewManager(db *gorm.DB) *ReviewManager {
	log.Info("[NewReviewManager] Initializing ReviewManager")

	return &ReviewManager{db}
}

func (m *ReviewManager) GetAll() ([]models.Review, *db_error.DBError) {
	log.Info("[ReviewManager.GetAll] Fetching all reviews")

	var allReviews []models.Review
	err := m.db.Find(&allReviews).Error
	if err != nil {
		log.Errorf("[ReviewManager.GetAll] Error fetching all reviews: %v", err)
		return nil, db_error.NewDBError(db_error.InternalError, "[ReviewManager.GetAll] Error fetching all reviews: %v", err)
	}

	log.Infof("[ReviewManager.GetAll] Successfully fetched all reviews")
	return allReviews, nil
}

func (m *ReviewManager) Get(idToGet uuid.UUID) (models.Review, *db_error.DBError) {
	log.Infof("[ReviewManager.Get] Fetching review with ID: %s", idToGet)

	var review models.Review
	err := m.db.First(&review, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[ReviewManager.Get] Error fetching review with ID %s: %v", idToGet, err)
		return models.Review{}, db_error.NewDBError(db_error.InternalError, "[ReviewManager.Get] Error fetching review with ID %s: %v", idToGet, err)
	}

	log.Infof("[ReviewManager.Get] Successfully fetched review with ID: %s", idToGet)
	return review, nil
}

func (m *ReviewManager) Create(newReview models.Review) (models.Review, *db_error.DBError) {
	log.Infof("[ReviewManager.Create] Creating new review")

	err := newReview.Validate()
	if err != nil {
		return models.Review{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	newReview.ID = uuid.New()

	err = m.db.Create(&newReview).Error
	if err != nil {
		log.Errorf("[ReviewManager.Create] Error creating new review with ID %s: %v", newReview.ID, err)
		return models.Review{}, db_error.NewDBError(db_error.InternalError, "[ReviewManager.Create] Error creating new review with ID %s: %v", newReview.ID, err)
	}

	log.Infof("[ReviewManager.Create] Successfully created review with ID: %s", newReview.ID)
	return newReview, nil
}

func (m *ReviewManager) Update(updatedReview models.Review) (models.Review, *db_error.DBError) {
	log.Infof("[ReviewManager.Update] Updating review with ID: %s", updatedReview.ID)

	err := updatedReview.Validate()
	if err != nil {
		return models.Review{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	var review models.Review
	err = m.db.First(&review, "id = ?", updatedReview.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[ReviewManager.Update] Review with ID %s does not exist", updatedReview.ID)
			return models.Review{}, db_error.NewDBError(db_error.NotFoundError, "[ReviewManager.Update] Review with ID %s does not exist", updatedReview.ID)
		}
		log.Errorf("[ReviewManager.Update] Error fetching review with ID %s: %v", updatedReview.ID, err)
		return models.Review{}, db_error.NewDBError(db_error.InternalError, "[ReviewManager.Update] Error fetching review with ID %s: %v", updatedReview.ID, err)
	}

	err = m.db.Model(&review).Updates(updatedReview).Error
	if err != nil {
		log.Errorf("[ReviewManager.Update] Error updating review with ID %s: %v", updatedReview.ID, err)
		return models.Review{}, db_error.NewDBError(db_error.InternalError, "[ReviewManager.Update] Error updating review with ID %s: %v", updatedReview.ID, err)
	}

	log.Infof("[ReviewManager.Update] Successfully updated review with ID: %s", updatedReview.ID)
	return updatedReview, nil
}

func (m *ReviewManager) Delete(idToDelete uuid.UUID) (models.Review, *db_error.DBError) {
	log.Infof("[ReviewManager.Delete] Deleting review with ID: %s", idToDelete)

	var review models.Review
	err := m.db.First(&review, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[ReviewManager.Delete] Review with ID %s does not exist", idToDelete)
			return models.Review{}, db_error.NewDBError(db_error.NotFoundError, "[ReviewManager.Delete] Review with ID %s does not exist", idToDelete)
		}
		log.Errorf("[ReviewManager.Delete] Error fetching review with ID %s: %v", idToDelete, err)
		return models.Review{}, db_error.NewDBError(db_error.InternalError, "[ReviewManager.Delete] Error fetching review with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&review).Error
	if err != nil {
		log.Errorf("[ReviewManager.Delete] Error deleting review with ID %s: %v", idToDelete, err)
		return models.Review{}, db_error.NewDBError(db_error.InternalError, "[ReviewManager.Delete] Error deleting review with ID %s: %v", idToDelete, err)
	}

	log.Infof("[ReviewManager.Delete] Successfully deleted review with ID: %s", idToDelete)
	return review, nil
}
