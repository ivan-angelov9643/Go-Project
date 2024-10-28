package managers

import (
	"awesomeProject/todo-app/structs"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type TagManager struct {
	tags map[uuid.UUID]structs.Tag
}

func NewTagManager() *TagManager {
	log.Info("[NewTagManager] Initializing TagManager")

	tagID1 := uuid.New()
	tagID2 := uuid.New()
	tagID3 := uuid.New()

	return &TagManager{
		tags: map[uuid.UUID]structs.Tag{
			tagID1: {ID: tagID1, Name: "Tag1"},
			tagID2: {ID: tagID2, Name: "Tag2"},
			tagID3: {ID: tagID3, Name: "Tag3"},
		},
	}
}

func (m *TagManager) GetAll() []*structs.Tag {
	log.Info("[TagManager.GetAll] Fetching all tags")

	allTags := make([]*structs.Tag, 0, len(m.tags))
	for _, tag := range m.tags {
		allTags = append(allTags, &tag)
	}

	return allTags
}

func (m *TagManager) Get(idToGet uuid.UUID) (*structs.Tag, error) {
	log.Infof("[TagManager.Get] Fetching tag with ID: %s", idToGet)

	tag, exists := m.tags[idToGet]
	if !exists {
		log.Errorf("[TagManager.Get] Tag with ID %s not found", idToGet)
		return nil, fmt.Errorf("[TagManager.Get] Tag with id %s not found", idToGet)
	}

	return &tag, nil
}

func (m *TagManager) Create(newTag structs.Tag) (structs.Tag, error) {
	log.Infof("[TagManager.Create] Creating new tag")

	if newTag.ID == uuid.Nil {
		newTag.ID = uuid.New()
	}

	_, exists := m.tags[newTag.ID]
	if exists {
		log.Errorf("[TagManager.Create] Tag with ID %s already exists", newTag.ID)
		return structs.Tag{}, fmt.Errorf("[TagManager.Create] Tag with id %s already exists", newTag.ID)
	}

	m.tags[newTag.ID] = newTag
	log.Infof("[TagManager.Create] Successfully created tag with ID: %s", newTag.ID)
	return newTag, nil
}

func (m *TagManager) Update(updatedTag structs.Tag) (structs.Tag, error) {
	log.Infof("[TagManager.Update] Updating tag with ID: %s", updatedTag.ID)

	_, exists := m.tags[updatedTag.ID]
	if !exists {
		log.Errorf("[TagManager.Update] Tag with ID %s not found", updatedTag.ID)
		return structs.Tag{}, fmt.Errorf("[TagManager.Update] Tag with id %s not found", updatedTag.ID)
	}

	m.tags[updatedTag.ID] = updatedTag
	log.Infof("[TagManager.Update] Successfully updated tag with ID: %s", updatedTag.ID)
	return updatedTag, nil
}

func (m *TagManager) Delete(idToDelete uuid.UUID) (structs.Tag, error) {
	log.Infof("[TagManager.Delete] Deleting tag with ID: %s", idToDelete)

	_, exists := m.tags[idToDelete]
	if !exists {
		log.Errorf("[TagManager.Delete] Tag with ID %s not found", idToDelete)
		return structs.Tag{}, fmt.Errorf("[TagManager.Delete] Tag with id %s not found", idToDelete)
	}

	deletedTag := m.tags[idToDelete]
	delete(m.tags, idToDelete)
	log.Infof("[TagManager.Delete] Successfully deleted tag with ID: %s", idToDelete)
	return deletedTag, nil
}
