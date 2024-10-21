package managers

import (
	"awesomeProject/todo-app/structs"
	"fmt"
	"time"
)

type ListManager struct {
	lists map[int]structs.List
}

func NewListManager() *ListManager {
	return &ListManager{
		lists: map[int]structs.List{
			1: {
				ID:           1,
				Name:         "List 1",
				Description:  "Description for list 1",
				CreationTime: time.Now(),
				Items: []structs.Item{
					{
						ID:           1,
						ListID:       1,
						Tittle:       "Item 1",
						Description:  "Description for item 1",
						Tags:         []structs.Tag{{ID: 1, Name: "Tag1"}, {ID: 2, Name: "Tag2"}},
						Completed:    false,
						CreationTime: time.Now(),
					},
				},
			},
			2: {
				ID:           2,
				Name:         "List 2",
				Description:  "Description for list 2",
				CreationTime: time.Now().Add(-time.Hour * 48),
				Items:        []structs.Item{},
			},
		},
	}
}

func (m *ListManager) GetAll() []*structs.List {
	allLists := make([]*structs.List, 0, len(m.lists))
	for _, list := range m.lists {
		allLists = append(allLists, &list)
	}
	return allLists
}

func (m *ListManager) Get(id int) (*structs.List, error) {
	list, exists := m.lists[id]
	if !exists {
		return nil, fmt.Errorf("list with id %d not found", id)
	}
	return &list, nil
}

func (m *ListManager) Create(list structs.List) error {
	_, exists := m.lists[list.ID]
	if exists {
		return fmt.Errorf("list with id %d already exists", list.ID)
	}
	m.lists[list.ID] = list
	return nil
}

func (m *ListManager) Update(id int, updated structs.List) error {
	_, exists := m.lists[id]
	if !exists {
		return fmt.Errorf("list with id %d not found", id)
	}
	m.lists[id] = updated
	return nil
}

func (m *ListManager) Delete(id int) error {
	_, exists := m.lists[id]
	if !exists {
		return fmt.Errorf("list with id %d not found", id)
	}
	delete(m.lists, id)
	return nil
}
