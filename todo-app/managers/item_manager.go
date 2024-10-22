package managers

import (
	"awesomeProject/todo-app/structs"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type ItemManager struct {
	items map[uuid.UUID]structs.Item
}

func NewItemManager() *ItemManager {
	log.Debug("Initializing ItemManager")

	itemID1 := uuid.New()
	itemID2 := uuid.New()
	listID1 := uuid.New()

	return &ItemManager{
		items: map[uuid.UUID]structs.Item{
			itemID1: {
				ID:          itemID1,
				ListID:      listID1,
				Tittle:      "Item 1",
				Description: "Description for item 1",
				Tags: []structs.Tag{
					{ID: uuid.New(), Name: "Tag1"},
					{ID: uuid.New(), Name: "Tag2"},
				},
				Completed:    false,
				CreationTime: time.Now(),
			},
			itemID2: {
				ID:          itemID2,
				ListID:      listID1,
				Tittle:      "Item 2",
				Description: "Description for item 2",
				Tags: []structs.Tag{
					{ID: uuid.New(), Name: "Tag3"},
				},
				Completed:    true,
				CreationTime: time.Now().Add(-time.Hour * 24),
			},
		},
	}
}

func (m *ItemManager) GetAll() []*structs.Item {
	log.Debug("Fetching all items")
	allItems := make([]*structs.Item, 0, len(m.items))
	for _, item := range m.items {
		allItems = append(allItems, &item)
	}
	return allItems
}

func (m *ItemManager) Get(idToGet uuid.UUID) (*structs.Item, error) {
	log.Debugf("Fetching item with ID: %s", idToGet)
	item, exists := m.items[idToGet]
	if !exists {
		log.Errorf("[Get] Item with ID %s not found", idToGet)
		return nil, fmt.Errorf("[Get] Item with id %s not found", idToGet)
	}
	return &item, nil
}

func (m *ItemManager) Create(newItem structs.Item) (structs.Item, error) {
	if newItem.ID == uuid.Nil {
		newItem.ID = uuid.New()
	}
	log.Debugf("Creating new item with ID: %s", newItem.ID)
	_, exists := m.items[newItem.ID]
	if exists {
		log.Errorf("[Create] Item with ID %s already exists", newItem.ID)
		return structs.Item{}, fmt.Errorf("[Create] Item with id %s already exists", newItem.ID)
	}
	m.items[newItem.ID] = newItem
	log.Debugf("Successfully created item with ID: %s", newItem.ID)
	return newItem, nil
}

func (m *ItemManager) Update(updatedItem structs.Item) (structs.Item, error) {
	log.Debugf("Updating item with ID: %s", updatedItem.ID)
	_, exists := m.items[updatedItem.ID]
	if !exists {
		log.Errorf("[Update] Item with ID %s not found", updatedItem.ID)
		return structs.Item{}, fmt.Errorf("[Update] Item with id %s not found", updatedItem.ID)
	}
	m.items[updatedItem.ID] = updatedItem
	log.Debugf("Successfully updated item with ID: %s", updatedItem.ID)
	return updatedItem, nil
}

func (m *ItemManager) Delete(idToDelete uuid.UUID) (structs.Item, error) {
	log.Debugf("Deleting item with ID: %s", idToDelete)
	_, exists := m.items[idToDelete]
	if !exists {
		log.Errorf("[Delete] Item with ID %s not found", idToDelete)
		return structs.Item{}, fmt.Errorf("[Delete] Item with id %s not found", idToDelete)
	}
	deletedItem := m.items[idToDelete]
	delete(m.items, idToDelete)
	log.Debugf("Successfully deleted item with ID: %s", idToDelete)
	return deletedItem, nil
}
