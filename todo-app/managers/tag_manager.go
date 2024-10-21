package managers

import (
	"awesomeProject/todo-app/structs"
	"fmt"
)

type TagManager struct {
	tags map[int]structs.Tag
}

func NewTagManager() *TagManager {
	return &TagManager{
		tags: map[int]structs.Tag{
			1: {ID: 1, Name: "Tag1"},
			2: {ID: 2, Name: "Tag2"},
			3: {ID: 3, Name: "Tag3"},
		},
	}
}

func (m *TagManager) GetAll() []*structs.Tag {
	allTags := make([]*structs.Tag, 0, len(m.tags))
	for _, tag := range m.tags {
		allTags = append(allTags, &tag)
	}
	return allTags
}

func (m *TagManager) Get(id int) (*structs.Tag, error) {
	tag, exists := m.tags[id]
	if !exists {
		return nil, fmt.Errorf("tag with id %d not found", id)
	}
	return &tag, nil
}

func (m *TagManager) Create(tag structs.Tag) error {
	_, exists := m.tags[tag.ID]
	if exists {
		return fmt.Errorf("tag with id %d already exists", tag.ID)
	}
	m.tags[tag.ID] = tag
	return nil
}

func (m *TagManager) Update(id int, updated structs.Tag) error {
	_, exists := m.tags[id]
	if !exists {
		return fmt.Errorf("tag with id %d not found", id)
	}
	m.tags[id] = updated
	return nil
}

func (m *TagManager) Delete(id int) error {
	_, exists := m.tags[id]
	if !exists {
		return fmt.Errorf("tag with id %d not found", id)
	}
	delete(m.tags, id)
	return nil
}
