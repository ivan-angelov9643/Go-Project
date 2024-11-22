package implementations

import (
	"awesomeProject/todo-app/models"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryManager struct {
	db         *gorm.DB
	categories map[uuid.UUID]models.Category
}

func NewCategoryManager(db *gorm.DB) *CategoryManager {
	log.Info("[NewCategoryManager] Initializing CategoryManager")

	return &CategoryManager{db, make(map[uuid.UUID]models.Category)}
}

func (m *CategoryManager) GetAll() []models.Category {
	log.Info("[CategoryManager.GetAll] Fetching all categories")

	allCategories := make([]models.Category, 0, len(m.categories))
	for _, category := range m.categories {
		allCategories = append(allCategories, category)
	}

	log.Infof("[CategoryManager.GetAll] Successfully fetched all categories")
	return allCategories
}

func (m *CategoryManager) Get(idToGet uuid.UUID) (models.Category, error) {
	log.Infof("[CategoryManager.Get] Fetching category with ID: %s", idToGet)

	category, exists := m.categories[idToGet]
	if !exists {
		log.Errorf("[CategoryManager.Get] Category with ID %s not found", idToGet)
		return models.Category{}, fmt.Errorf("[CategoryManager.Get] Category with ID %s not found", idToGet)
	}

	log.Infof("[CategoryManager.Get] Successfully fetched category with ID: %s", idToGet)
	return category, nil
}

func (m *CategoryManager) Create(newCategory models.Category) (models.Category, error) {
	log.Infof("[CategoryManager.Create] Creating new category")

	if newCategory.ID == uuid.Nil {
		newCategory.ID = uuid.New()
	}

	_, exists := m.categories[newCategory.ID]
	if exists {
		log.Errorf("[CategoryManager.Create] Category with ID %s already exists", newCategory.ID)
		return models.Category{}, fmt.Errorf("[CategoryManager.Create] Category with ID %s already exists", newCategory.ID)
	}

	m.categories[newCategory.ID] = newCategory
	log.Infof("[CategoryManager.Create] Successfully created category with ID: %s", newCategory.ID)
	return newCategory, nil
}

func (m *CategoryManager) Update(updatedCategory models.Category) (models.Category, error) {
	log.Infof("[CategoryManager.Update] Updating category with ID: %s", updatedCategory.ID)

	_, exists := m.categories[updatedCategory.ID]
	if !exists {
		log.Errorf("[CategoryManager.Update] Category with ID %s not found", updatedCategory.ID)
		return models.Category{}, fmt.Errorf("[CategoryManager.Update] Category with ID %s not found", updatedCategory.ID)
	}

	m.categories[updatedCategory.ID] = updatedCategory
	log.Infof("[CategoryManager.Update] Successfully updated category with ID: %s", updatedCategory.ID)
	return updatedCategory, nil
}

func (m *CategoryManager) Delete(idToDelete uuid.UUID) (models.Category, error) {
	log.Infof("[CategoryManager.Delete] Deleting category with ID: %s", idToDelete)

	deletedCategory, exists := m.categories[idToDelete]
	if !exists {
		log.Errorf("[CategoryManager.Delete] Category with ID %s not found", idToDelete)
		return models.Category{}, fmt.Errorf("[CategoryManager.Delete] Category with ID %s not found", idToDelete)
	}

	delete(m.categories, idToDelete)
	log.Infof("[CategoryManager.Delete] Successfully deleted category with ID: %s", idToDelete)
	return deletedCategory, nil
}
