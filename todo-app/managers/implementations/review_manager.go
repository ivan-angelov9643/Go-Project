package implementations

import (
	"awesomeProject/todo-app/models"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReviewManager struct {
	db      *gorm.DB
	reviews map[uuid.UUID]models.Review
}

func NewReviewManager(db *gorm.DB) *ReviewManager {
	log.Info("[NewReviewManager] Initializing ReviewManager")

	return &ReviewManager{db, make(map[uuid.UUID]models.Review)}
}

func (m *ReviewManager) GetAll() []models.Review {
	log.Info("[ReviewManager.GetAll] Fetching all reviews")

	allReviews := make([]models.Review, 0, len(m.reviews))
	for _, review := range m.reviews {
		allReviews = append(allReviews, review)
	}

	log.Infof("[ReviewManager.GetAll] Successfully fetched all reviews")
	return allReviews
}

func (m *ReviewManager) Get(idToGet uuid.UUID) (models.Review, error) {
	log.Infof("[ReviewManager.Get] Fetching review with ID: %s", idToGet)

	review, exists := m.reviews[idToGet]
	if !exists {
		log.Errorf("[ReviewManager.Get] Review with ID %s not found", idToGet)
		return models.Review{}, fmt.Errorf("[ReviewManager.Get] Review with ID %s not found", idToGet)
	}

	log.Infof("[ReviewManager.Get] Successfully fetched review with ID: %s", idToGet)
	return review, nil
}

func (m *ReviewManager) Create(newReview models.Review) (models.Review, error) {
	log.Infof("[ReviewManager.Create] Creating new review")

	if newReview.ID == uuid.Nil {
		newReview.ID = uuid.New()
	}

	_, exists := m.reviews[newReview.ID]
	if exists {
		log.Errorf("[ReviewManager.Create] Review with ID %s already exists", newReview.ID)
		return models.Review{}, fmt.Errorf("[ReviewManager.Create] Review with ID %s already exists", newReview.ID)
	}

	m.reviews[newReview.ID] = newReview
	log.Infof("[ReviewManager.Create] Successfully created review with ID: %s", newReview.ID)
	return newReview, nil
}

func (m *ReviewManager) Update(updatedReview models.Review) (models.Review, error) {
	log.Infof("[ReviewManager.Update] Updating review with ID: %s", updatedReview.ID)

	_, exists := m.reviews[updatedReview.ID]
	if !exists {
		log.Errorf("[ReviewManager.Update] Review with ID %s not found", updatedReview.ID)
		return models.Review{}, fmt.Errorf("[ReviewManager.Update] Review with ID %s not found", updatedReview.ID)
	}

	m.reviews[updatedReview.ID] = updatedReview
	log.Infof("[ReviewManager.Update] Successfully updated review with ID: %s", updatedReview.ID)
	return updatedReview, nil
}

func (m *ReviewManager) Delete(idToDelete uuid.UUID) (models.Review, error) {
	log.Infof("[ReviewManager.Delete] Deleting review with ID: %s", idToDelete)

	deletedReview, exists := m.reviews[idToDelete]
	if !exists {
		log.Errorf("[ReviewManager.Delete] Review with ID %s not found", idToDelete)
		return models.Review{}, fmt.Errorf("[ReviewManager.Delete] Review with ID %s not found", idToDelete)
	}

	delete(m.reviews, idToDelete)
	log.Infof("[ReviewManager.Delete] Successfully deleted review with ID: %s", idToDelete)
	return deletedReview, nil
}
