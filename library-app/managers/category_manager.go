package managers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryManager struct {
	db *gorm.DB
}

func NewCategoryManager(db *gorm.DB) *CategoryManager {
	log.Info("[NewCategoryManager] Initializing CategoryManager")

	return &CategoryManager{db}
}

func (m *CategoryManager) GetAll(scopes ...db.DBScope) ([]models.Category, error) {
	log.Info("[CategoryManager.GetAll] Fetching all categories")

	var allCategories []models.Category
	err := db.ApplyScopes(m.db, scopes).Find(&allCategories).Error
	if err != nil {
		log.Errorf("[CategoryManager.GetAll] Error fetching all categories: %v", err)
		return nil, db.NewDBError(db.InternalError, "[CategoryManager.GetAll] Error fetching all categories: %v", err)
	}

	log.Infof("[CategoryManager.GetAll] Successfully fetched all categories")
	return allCategories, nil
}

func (m *CategoryManager) Get(idToGet uuid.UUID) (models.Category, error) {
	log.Infof("[CategoryManager.Get] Fetching category with ID: %s", idToGet)

	var category models.Category
	err := m.db.First(&category, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[CategoryManager.Get] Error fetching category with ID %s: %v", idToGet, err)
		return models.Category{}, db.NewDBError(db.InternalError, "[CategoryManager.Get] Error fetching category with ID %s: %v", idToGet, err)
	}

	log.Infof("[CategoryManager.Get] Successfully fetched category with ID: %s", idToGet)
	return category, nil
}

func (m *CategoryManager) Create(newCategory models.Category) (models.Category, error) {
	log.Infof("[CategoryManager.Create] Creating new category")

	err := newCategory.Validate()
	if err != nil {
		return models.Category{}, db.NewDBError(db.ValidationError, err.Error())
	}

	newCategory.ID = uuid.New()

	err = m.db.Create(&newCategory).Error
	if err != nil {
		log.Errorf("[CategoryManager.Create] Error creating new category with ID %s: %v", newCategory.ID, err)
		return models.Category{}, db.NewDBError(db.InternalError, "[CategoryManager.Create] Error creating new category with ID %s: %v", newCategory.ID, err)
	}

	log.Infof("[CategoryManager.Create] Successfully created category with ID: %s", newCategory.ID)
	return newCategory, nil
}

func (m *CategoryManager) Update(updatedCategory models.Category) (models.Category, error) {
	log.Infof("[CategoryManager.Update] Updating category with ID: %s", updatedCategory.ID)

	err := updatedCategory.Validate()
	if err != nil {
		return models.Category{}, db.NewDBError(db.ValidationError, err.Error())
	}

	var category models.Category
	err = m.db.First(&category, "id = ?", updatedCategory.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[CategoryManager.Update] Category with ID %s does not exist", updatedCategory.ID)
			return models.Category{}, db.NewDBError(db.NotFoundError, "[CategoryManager.Update] Category with ID %s does not exist", updatedCategory.ID)
		}
		log.Errorf("[CategoryManager.Update] Error fetching category with ID %s: %v", updatedCategory.ID, err)
		return models.Category{}, db.NewDBError(db.InternalError, "[CategoryManager.Update] Error fetching category with ID %s: %v", updatedCategory.ID, err)
	}

	err = m.db.Model(&category).Updates(updatedCategory).Error
	if err != nil {
		log.Errorf("[CategoryManager.Update] Error updating category with ID %s: %v", updatedCategory.ID, err)
		return models.Category{}, db.NewDBError(db.InternalError, "[CategoryManager.Update] Error updating category with ID %s: %v", updatedCategory.ID, err)
	}

	log.Infof("[CategoryManager.Update] Successfully updated category with ID: %s", updatedCategory.ID)
	return updatedCategory, nil
}

func (m *CategoryManager) Delete(idToDelete uuid.UUID) (models.Category, error) {
	log.Infof("[CategoryManager.Delete] Deleting category with ID: %s", idToDelete)

	var category models.Category
	err := m.db.First(&category, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[CategoryManager.Delete] Category with ID %s does not exist", idToDelete)
			return models.Category{}, db.NewDBError(db.NotFoundError, "[CategoryManager.Delete] Category with ID %s does not exist", idToDelete)
		}
		log.Errorf("[CategoryManager.Delete] Error fetching category with ID %s: %v", idToDelete, err)
		return models.Category{}, db.NewDBError(db.InternalError, "[CategoryManager.Delete] Error fetching category with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&category).Error
	if err != nil {
		log.Errorf("[CategoryManager.Delete] Error deleting category with ID %s: %v", idToDelete, err)
		return models.Category{}, db.NewDBError(db.InternalError, "[CategoryManager.Delete] Error deleting category with ID %s: %v", idToDelete, err)
	}

	log.Infof("[CategoryManager.Delete] Successfully deleted category with ID: %s", idToDelete)
	return category, nil
}

func (m *CategoryManager) Count(scopes ...db.DBScope) (int64, error) {
	log.Infof("[CategoryManager.Count] Counting categories in the database")

	var count int64
	err := db.ApplyScopes(m.db, scopes).Model(&models.Category{}).Count(&count).Error
	if err != nil {
		log.Errorf("[CategoryManager.Count] Error counting categories: %v", err)
		return 0, db.NewDBError(db.InternalError, "[CategoryManager.Count] Error counting categories: %v", err)
	}

	log.Infof("[CategoryManager.Count] Successfully counted categories: %d", count)
	return count, nil
}
