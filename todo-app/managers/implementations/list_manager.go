package implementations

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
	log.Info("[NewListManager] Initializing ListManager")

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
						Title:       "Item 1",
						Description: "Description for item 1",
						Tags: []structs.Tag{
							{Name: "Tag1"},
							{Name: "Tag2"},
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

func (m *ListManager) GetAll() []structs.List {
	log.Info("[ListManager.GetAll] Fetching all lists")

	allLists := make([]structs.List, 0, len(m.lists))
	for _, list := range m.lists {
		allLists = append(allLists, list)
	}

	return allLists
}

func (m *ListManager) Get(idToGet uuid.UUID) (structs.List, error) {
	log.Infof("[ListManager.Get] Fetching list with ID: %s", idToGet)

	list, exists := m.lists[idToGet]
	if !exists {
		log.Errorf("[ListManager.Get] List with ID %s not found", idToGet)
		return structs.List{}, fmt.Errorf("[ListManager.Get] List with id %s not found", idToGet)
	}

	return list, nil
}

func (m *ListManager) Create(newList structs.List) (structs.List, error) {
	log.Info("[ListManager.Create] Creating new list")

	if newList.ID == uuid.Nil {
		newList.ID = uuid.New()
	}

	_, exists := m.lists[newList.ID]
	if exists {
		log.Errorf("[ListManager.Create] List with ID %s already exists", newList.ID)
		return structs.List{}, fmt.Errorf("[ListManager.Create] List with id %s already exists", newList.ID)
	}

	m.lists[newList.ID] = newList
	log.Infof("[ListManager.Create] Successfully created list with ID: %s", newList.ID)
	return newList, nil
}

func (m *ListManager) Update(updatedList structs.List) (structs.List, error) {
	log.Infof("[ListManager.Update] Updating list with ID: %s", updatedList.ID)

	_, exists := m.lists[updatedList.ID]
	if !exists {
		log.Errorf("[ListManager.Update] List with ID %s not found", updatedList.ID)
		return structs.List{}, fmt.Errorf("[ListManager.Update] List with id %s not found", updatedList.ID)
	}

	m.lists[updatedList.ID] = updatedList
	log.Infof("[ListManager.Update] Successfully updated list with ID: %s", updatedList.ID)
	return updatedList, nil
}

func (m *ListManager) Delete(idToDelete uuid.UUID) (structs.List, error) {
	log.Infof("[ListManager.Delete] Deleting list with ID: %s", idToDelete)

	_, exists := m.lists[idToDelete]
	if !exists {
		log.Errorf("[ListManager.Delete] List with ID %s not found", idToDelete)
		return structs.List{}, fmt.Errorf("[ListManager.Delete] List with id %s not found", idToDelete)
	}

	deletedList := m.lists[idToDelete]
	delete(m.lists, idToDelete)
	log.Infof("[ListManager.Delete] Successfully deleted list with ID: %s", idToDelete)
	return deletedList, nil
}
