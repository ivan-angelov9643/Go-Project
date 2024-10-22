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

func (m *ItemManager) Get(id uuid.UUID) (*structs.Item, error) {
	log.Debugf("Fetching item with ID: %s", id)
	item, exists := m.items[id]
	if !exists {
		log.Errorf("Item with ID %s not found", id)
		return nil, fmt.Errorf("item with id %s not found", id)
	}
	return &item, nil
}

func (m *ItemManager) Create(item structs.Item) error {
	log.Debugf("Creating new item with ID: %s", item.ID)
	_, exists := m.items[item.ID]
	if exists {
		log.Errorf("Item with ID %s already exists", item.ID)
		return fmt.Errorf("item with id %s already exists", item.ID)
	}
	m.items[item.ID] = item
	log.Debugf("Successfully created item with ID: %s", item.ID)
	return nil
}

func (m *ItemManager) Update(id uuid.UUID, updated structs.Item) error {
	log.Debugf("Updating item with ID: %s", id)
	_, exists := m.items[id]
	if !exists {
		log.Errorf("Item with ID %s not found", id)
		return fmt.Errorf("item with id %s not found", id)
	}
	m.items[id] = updated
	log.Debugf("Successfully updated item with ID: %s", id)
	return nil
}

func (m *ItemManager) Delete(id uuid.UUID) error {
	log.Debugf("Deleting item with ID: %s", id)
	_, exists := m.items[id]
	if !exists {
		log.Errorf("Item with ID %s not found", id)
		return fmt.Errorf("item with id %s not found", id)
	}
	delete(m.items, id)
	log.Debugf("Successfully deleted item with ID: %s", id)
	return nil
}
