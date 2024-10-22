package managers

import (
	"awesomeProject/todo-app/structs"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type ListManager struct {
	lists map[uuid.UUID]structs.List
}

func NewListManager() *ListManager {
	log.Debug("Initializing ListManager")

	listID1 := uuid.New()
	listID2 := uuid.New()
	itemID1 := uuid.New()

	return &ListManager{
		lists: map[uuid.UUID]structs.List{
			listID1: {
				ID:           listID1,
				Name:         "List 1",
				Description:  "Description for list 1",
				CreationTime: time.Now(),
				Items: []structs.Item{
					{
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
				},
			},
			listID2: {
				ID:           listID2,
				Name:         "List 2",
				Description:  "Description for list 2",
				CreationTime: time.Now().Add(-time.Hour * 48),
				Items:        []structs.Item{},
			},
		},
	}
}

func (m *ListManager) GetAll() []*structs.List {
	log.Debug("Fetching all lists")
	allLists := make([]*structs.List, 0, len(m.lists))
	for _, list := range m.lists {
		allLists = append(allLists, &list)
	}
	return allLists
}

func (m *ListManager) Get(id uuid.UUID) (*structs.List, error) {
	log.Debugf("Fetching list with ID: %s", id)
	list, exists := m.lists[id]
	if !exists {
		log.Errorf("List with ID %s not found", id)
		return nil, fmt.Errorf("list with id %s not found", id)
	}
	return &list, nil
}

func (m *ListManager) Create(list structs.List) error {
	log.Debugf("Creating new list with ID: %s", list.ID)
	_, exists := m.lists[list.ID]
	if exists {
		log.Errorf("List with ID %s already exists", list.ID)
		return fmt.Errorf("list with id %s already exists", list.ID)
	}
	m.lists[list.ID] = list
	log.Debugf("Successfully created list with ID: %s", list.ID)
	return nil
}

func (m *ListManager) Update(id uuid.UUID, updated structs.List) error {
	log.Debugf("Updating list with ID: %s", id)
	_, exists := m.lists[id]
	if !exists {
		log.Errorf("List with ID %s not found", id)
		return fmt.Errorf("list with id %s not found", id)
	}
	m.lists[id] = updated
	log.Debugf("Successfully updated list with ID: %s", id)
	return nil
}

func (m *ListManager) Delete(id uuid.UUID) error {
	log.Debugf("Deleting list with ID: %s", id)
	_, exists := m.lists[id]
	if !exists {
		log.Errorf("List with ID %s not found", id)
		return fmt.Errorf("list with id %s not found", id)
	}
	delete(m.lists, id)
	log.Debugf("Successfully deleted list with ID: %s", id)
	return nil
}
