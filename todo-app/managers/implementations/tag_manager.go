package implementations

import (
	"awesomeProject/todo-app/structs"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type TagManager struct {
	tags map[string]structs.Tag
}

func NewTagManager() *TagManager {
	log.Info("[NewTagManager] Initializing TagManager")

	return &TagManager{
		tags: map[string]structs.Tag{
			"Tag1": {Name: "Tag1"},
			"Tag2": {Name: "Tag2"},
			"Tag3": {Name: "Tag3"},
		},
	}
}

func (m *TagManager) GetAll() []structs.Tag {
	log.Info("[TagManager.GetAll] Fetching all tags")

	allTags := make([]structs.Tag, 0, len(m.tags))
	for _, tag := range m.tags {
		allTags = append(allTags, tag)
	}

	return allTags
}

func (m *TagManager) Get(nameToGet string) (structs.Tag, error) {
	log.Infof("[TagManager.Get] Fetching tag with name: %s", nameToGet)

	tag, exists := m.tags[nameToGet]
	if !exists {
		log.Errorf("[TagManager.Get] Tag with name %s not found", nameToGet)
		return structs.Tag{}, fmt.Errorf("[TagManager.Get] Tag with name %s not found", nameToGet)
	}

	return tag, nil
}

func (m *TagManager) Create(newTag structs.Tag) (structs.Tag, error) {
	log.Infof("[TagManager.Create] Creating new tag")

	_, exists := m.tags[newTag.Name]
	if exists {
		log.Errorf("[TagManager.Create] Tag with name %s already exists", newTag.Name)
		return structs.Tag{}, fmt.Errorf("[TagManager.Create] Tag with name %s already exists", newTag.Name)
	}

	m.tags[newTag.Name] = newTag
	log.Infof("[TagManager.Create] Successfully created tag with name: %s", newTag.Name)
	return newTag, nil
}

func (m *TagManager) Delete(nameToDelete string) (structs.Tag, error) {
	log.Infof("[TagManager.Delete] Deleting tag with name: %s", nameToDelete)

	_, exists := m.tags[nameToDelete]
	if !exists {
		log.Errorf("[TagManager.Delete] Tag with name %s not found", nameToDelete)
		return structs.Tag{}, fmt.Errorf("[TagManager.Delete] Tag with name %s not found", nameToDelete)
	}

	deletedTag := m.tags[nameToDelete]
	delete(m.tags, nameToDelete)
	log.Infof("[TagManager.Delete] Successfully deleted tag with name: %s", nameToDelete)
	return deletedTag, nil
}
