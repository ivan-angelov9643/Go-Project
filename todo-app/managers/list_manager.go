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

func (m *ListManager) Get(idToGet uuid.UUID) (*structs.List, error) {
	log.Debugf("Fetching list with ID: %s", idToGet)
	list, exists := m.lists[idToGet]
	if !exists {
		log.Errorf("[Get] List with ID %s not found", idToGet)
		return nil, fmt.Errorf("[Get] List with id %s not found", idToGet)
	}
	return &list, nil
}

func (m *ListManager) Create(newList structs.List) (structs.List, error) {
	if newList.ID == uuid.Nil {
		newList.ID = uuid.New()
	}
	log.Debugf("Creating new list with ID: %s", newList.ID)
	_, exists := m.lists[newList.ID]
	if exists {
		log.Errorf("[Create] List with ID %s already exists", newList.ID)
		return structs.List{}, fmt.Errorf("[Create] List with id %s already exists", newList.ID)
	}
	m.lists[newList.ID] = newList
	log.Debugf("Successfully created list with ID: %s", newList.ID)
	return newList, nil
}

func (m *ListManager) Update(updatedList structs.List) (structs.List, error) {
	log.Debugf("Updating list with ID: %s", updatedList.ID)
	_, exists := m.lists[updatedList.ID]
	if !exists {
		log.Errorf("[Update] List with ID %s not found", updatedList.ID)
		return structs.List{}, fmt.Errorf("[Update] List with id %s not found", updatedList.ID)
	}
	m.lists[updatedList.ID] = updatedList
	log.Debugf("Successfully updated list with ID: %s", updatedList.ID)
	return updatedList, nil
}

func (m *ListManager) Delete(idToDelete uuid.UUID) (structs.List, error) {
	log.Debugf("Deleting list with ID: %s", idToDelete)
	_, exists := m.lists[idToDelete]
	if !exists {
		log.Errorf("[Delete] List with ID %s not found", idToDelete)
		return structs.List{}, fmt.Errorf("[Delete] List with id %s not found", idToDelete)
	}
	deletedList := m.lists[idToDelete]
	delete(m.lists, idToDelete)
	log.Debugf("Successfully deleted list with ID: %s", idToDelete)
	return deletedList, nil
}
