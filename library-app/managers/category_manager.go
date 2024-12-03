package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"errors"
	"github.com/google/uuid"
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

func (m *CategoryManager) GetAll() ([]models.Category, *db_error.DBError) {
	log.Info("[CategoryManager.GetAll] Fetching all categories")

	var allCategories []models.Category
	err := m.db.Find(&allCategories).Error
	if err != nil {
		log.Errorf("[CategoryManager.GetAll] Error fetching all categories: %v", err)
		return nil, db_error.NewDBError(db_error.InternalError, "[CategoryManager.GetAll] Error fetching all categories: %v", err)
	}

	log.Infof("[CategoryManager.GetAll] Successfully fetched all categories")
	return allCategories, nil
}

func (m *CategoryManager) Get(idToGet uuid.UUID) (models.Category, *db_error.DBError) {
	log.Infof("[CategoryManager.Get] Fetching category with ID: %s", idToGet)

	var category models.Category
	err := m.db.First(&category, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[CategoryManager.Get] Error fetching category with ID %s: %v", idToGet, err)
		return models.Category{}, db_error.NewDBError(db_error.InternalError, "[CategoryManager.Get] Error fetching category with ID %s: %v", idToGet, err)
	}

	log.Infof("[CategoryManager.Get] Successfully fetched category with ID: %s", idToGet)
	return category, nil
}

func (m *CategoryManager) Create(newCategory models.Category) (models.Category, *db_error.DBError) {
	log.Infof("[CategoryManager.Create] Creating new category")

	err := newCategory.Validate()
	if err != nil {
		return models.Category{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	newCategory.ID = uuid.New()

	err = m.db.Create(&newCategory).Error
	if err != nil {
		log.Errorf("[CategoryManager.Create] Error creating new category with ID %s: %v", newCategory.ID, err)
		return models.Category{}, db_error.NewDBError(db_error.InternalError, "[CategoryManager.Create] Error creating new category with ID %s: %v", newCategory.ID, err)
	}

	log.Infof("[CategoryManager.Create] Successfully created category with ID: %s", newCategory.ID)
	return newCategory, nil
}

func (m *CategoryManager) Update(updatedCategory models.Category) (models.Category, *db_error.DBError) {
	log.Infof("[CategoryManager.Update] Updating category with ID: %s", updatedCategory.ID)

	err := updatedCategory.Validate()
	if err != nil {
		return models.Category{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	var category models.Category
	err = m.db.First(&category, "id = ?", updatedCategory.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[CategoryManager.Update] Category with ID %s does not exist", updatedCategory.ID)
			return models.Category{}, db_error.NewDBError(db_error.NotFoundError, "[CategoryManager.Update] Category with ID %s does not exist", updatedCategory.ID)
		}
		log.Errorf("[CategoryManager.Update] Error fetching category with ID %s: %v", updatedCategory.ID, err)
		return models.Category{}, db_error.NewDBError(db_error.InternalError, "[CategoryManager.Update] Error fetching category with ID %s: %v", updatedCategory.ID, err)
	}

	err = m.db.Model(&category).Updates(updatedCategory).Error
	if err != nil {
		log.Errorf("[CategoryManager.Update] Error updating category with ID %s: %v", updatedCategory.ID, err)
		return models.Category{}, db_error.NewDBError(db_error.InternalError, "[CategoryManager.Update] Error updating category with ID %s: %v", updatedCategory.ID, err)
	}

	log.Infof("[CategoryManager.Update] Successfully updated category with ID: %s", updatedCategory.ID)
	return updatedCategory, nil
}

func (m *CategoryManager) Delete(idToDelete uuid.UUID) (models.Category, *db_error.DBError) {
	log.Infof("[CategoryManager.Delete] Deleting category with ID: %s", idToDelete)

	var category models.Category
	err := m.db.First(&category, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[CategoryManager.Delete] Category with ID %s does not exist", idToDelete)
			return models.Category{}, db_error.NewDBError(db_error.NotFoundError, "[CategoryManager.Delete] Category with ID %s does not exist", idToDelete)
		}
		log.Errorf("[CategoryManager.Delete] Error fetching category with ID %s: %v", idToDelete, err)
		return models.Category{}, db_error.NewDBError(db_error.InternalError, "[CategoryManager.Delete] Error fetching category with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&category).Error
	if err != nil {
		log.Errorf("[CategoryManager.Delete] Error deleting category with ID %s: %v", idToDelete, err)
		return models.Category{}, db_error.NewDBError(db_error.InternalError, "[CategoryManager.Delete] Error deleting category with ID %s: %v", idToDelete, err)
	}

	log.Infof("[CategoryManager.Delete] Successfully deleted category with ID: %s", idToDelete)
	return category, nil
}
