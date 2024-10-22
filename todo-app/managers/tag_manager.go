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
	log.Debug("Initializing TagManager")

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
	log.Debug("Fetching all tags")
	allTags := make([]*structs.Tag, 0, len(m.tags))
	for _, tag := range m.tags {
		allTags = append(allTags, &tag)
	}
	return allTags
}

func (m *TagManager) Get(id uuid.UUID) (*structs.Tag, error) {
	log.Debugf("Fetching tag with ID: %s", id)
	tag, exists := m.tags[id]
	if !exists {
		log.Errorf("Tag with ID %s not found", id)
		return nil, fmt.Errorf("tag with id %s not found", id)
	}
	return &tag, nil
}

func (m *TagManager) Create(tag structs.Tag) error {
	log.Debugf("Creating new tag with ID: %s", tag.ID)
	_, exists := m.tags[tag.ID]
	if exists {
		log.Errorf("Tag with ID %s already exists", tag.ID)
		return fmt.Errorf("tag with id %s already exists", tag.ID)
	}
	m.tags[tag.ID] = tag
	log.Debugf("Successfully created tag with ID: %s", tag.ID)
	return nil
}

func (m *TagManager) Update(id uuid.UUID, updated structs.Tag) error {
	log.Debugf("Updating tag with ID: %s", id)
	_, exists := m.tags[id]
	if !exists {
		log.Errorf("Tag with ID %s not found", id)
		return fmt.Errorf("tag with id %s not found", id)
	}
	m.tags[id] = updated
	log.Debugf("Successfully updated tag with ID: %s", id)
	return nil
}

func (m *TagManager) Delete(id uuid.UUID) error {
	log.Debugf("Deleting tag with ID: %s", id)
	_, exists := m.tags[id]
	if !exists {
		log.Errorf("Tag with ID %s not found", id)
		return fmt.Errorf("tag with id %s not found", id)
	}
	delete(m.tags, id)
	log.Debugf("Successfully deleted tag with ID: %s", id)
	return nil
}
