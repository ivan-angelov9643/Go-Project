package managers

import (
	"awesomeProject/todo-app/structs"
	"fmt"
	"time"
)

type ItemManager struct {
	items map[int]structs.Item
}

func NewItemManager() *ItemManager {
	return &ItemManager{
		items: map[int]structs.Item{
			1: {
				ID:          1,
				ListID:      1,
				Tittle:      "Item 1",
				Description: "Description for item 1",
				Tags: []structs.Tag{
					{ID: 1, Name: "Tag1"},
					{ID: 2, Name: "Tag2"},
				},
				Completed:    false,
				CreationTime: time.Now(),
			},
			2: {
				ID:          2,
				ListID:      1,
				Tittle:      "Item 2",
				Description: "Description for item 2",
				Tags: []structs.Tag{
					{ID: 3, Name: "Tag3"},
				},
				Completed:    true,
				CreationTime: time.Now().Add(-time.Hour * 24),
			},
		},
	}
}

func (m *ItemManager) GetAll() []*structs.Item {
	allItems := make([]*structs.Item, 0, len(m.items))
	for _, item := range m.items {
		allItems = append(allItems, &item)
	}
	return allItems
}

func (m *ItemManager) Get(id int) (*structs.Item, error) {
	item, exists := m.items[id]
	if !exists {
		return nil, fmt.Errorf("item with id %d not found", id)
	}
	return &item, nil
}

func (m *ItemManager) Create(item structs.Item) error {
	_, exists := m.items[item.ID]
	if exists {
		return fmt.Errorf("item with id %d already exists", item.ID)
	}
	m.items[item.ID] = item
	return nil
}

func (m *ItemManager) Update(id int, updated structs.Item) error {
	_, exists := m.items[id]
	if !exists {
		//log
		return fmt.Errorf("item with id %d not found", id)
	}
	m.items[id] = updated
	return nil
}

func (m *ItemManager) Delete(id int) error {
	_, exists := m.items[id]
	if !exists {
		return fmt.Errorf("item with id %d not found", id)
	}
	delete(m.items, id)
	return nil
}
