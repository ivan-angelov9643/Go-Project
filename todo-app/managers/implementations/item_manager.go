package implementations

import (
	"awesomeProject/todo-app/models"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type ItemManager struct {
	items map[uuid.UUID]models.Item
}

func NewItemManager() *ItemManager {
	log.Info("[NewItemManager] Initializing ItemManager")

	itemID1 := uuid.New()
	itemID2 := uuid.New()
	listID1 := uuid.New()

	return &ItemManager{
		items: map[uuid.UUID]models.Item{
			itemID1: {
				ID:          itemID1,
				ListID:      listID1,
				Title:       "Item 1",
				Description: "Description for item 1",
				Tags: []models.Tag{
					{Name: "Tag1"},
					{Name: "Tag2"},
				},
				Completed:    false,
				CreationTime: time.Now(),
			},
			itemID2: {
				ID:          itemID2,
				ListID:      listID1,
				Title:       "Item 2",
				Description: "Description for item 2",
				Tags: []models.Tag{
					{Name: "Tag3"},
				},
				Completed:    true,
				CreationTime: time.Now().Add(-time.Hour * 24),
			},
		},
	}
}

func (m *ItemManager) GetAll() []models.Item {
	log.Info("[ItemManager.GetAll] Fetching all items")

	allItems := make([]models.Item, 0, len(m.items))
	for _, item := range m.items {
		allItems = append(allItems, item)
	}

	return allItems
}

func (m *ItemManager) Get(idToGet uuid.UUID) (models.Item, error) {
	log.Infof("[ItemManager.Get] Fetching item with ID: %s", idToGet)

	item, exists := m.items[idToGet]
	if !exists {
		log.Errorf("[ItemManager.Get] Item with ID %s not found", idToGet)
		return models.Item{}, fmt.Errorf("[ItemManager.Get] Item with id %s not found", idToGet)
	}

	return item, nil
}

func (m *ItemManager) Create(newItem models.Item) (models.Item, error) {
	log.Infof("[ItemManager.Create] Creating new item")

	if newItem.ID == uuid.Nil {
		newItem.ID = uuid.New()
	}

	_, exists := m.items[newItem.ID]
	if exists {
		log.Errorf("[ItemManager.Create] Item with ID %s already exists", newItem.ID)
		return models.Item{}, fmt.Errorf("[ItemManager.Create] Item with id %s already exists", newItem.ID)
	}

	m.items[newItem.ID] = newItem
	log.Infof("[ItemManager.Create] Successfully created item with ID: %s", newItem.ID)
	return newItem, nil
}

func (m *ItemManager) Update(updatedItem models.Item) (models.Item, error) {
	log.Infof("[ItemManager.Update] Updating item with ID: %s", updatedItem.ID)

	_, exists := m.items[updatedItem.ID]
	if !exists {
		log.Errorf("[ItemManager.Update] Item with ID %s not found", updatedItem.ID)
		return models.Item{}, fmt.Errorf("[ItemManager.Update] Item with id %s not found", updatedItem.ID)
	}

	m.items[updatedItem.ID] = updatedItem
	log.Infof("[ItemManager.Update] Successfully updated item with ID: %s", updatedItem.ID)
	return updatedItem, nil
}

func (m *ItemManager) Delete(idToDelete uuid.UUID) (models.Item, error) {
	log.Infof("[ItemManager.Delete] Deleting item with ID: %s", idToDelete)

	_, exists := m.items[idToDelete]
	if !exists {
		log.Errorf("[ItemManager.Delete] Item with ID %s not found", idToDelete)
		return models.Item{}, fmt.Errorf("[ItemManager.Delete] Item with id %s not found", idToDelete)
	}

	deletedItem := m.items[idToDelete]
	delete(m.items, idToDelete)
	log.Infof("[ItemManager.Delete] Successfully deleted item with ID: %s", idToDelete)
	return deletedItem, nil
}
