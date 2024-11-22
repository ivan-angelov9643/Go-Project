package implementations

import (
	"awesomeProject/todo-app/models"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthorManager struct {
	db      *gorm.DB
	authors map[uuid.UUID]models.Author
}

func NewAuthorManager(db *gorm.DB) *AuthorManager {
	log.Info("[NewAuthorManager] Initializing AuthorManager")

	return &AuthorManager{db, make(map[uuid.UUID]models.Author)}
}

func (m *AuthorManager) GetAll() []models.Author {
	log.Info("[AuthorManager.GetAll] Fetching all authors")

	allAuthors := make([]models.Author, 0, len(m.authors))
	for _, author := range m.authors {
		allAuthors = append(allAuthors, author)
	}

	log.Infof("[AuthorManager.GetAll] Successfully fetched all authors")
	return allAuthors
}

func (m *AuthorManager) Get(idToGet uuid.UUID) (models.Author, error) {
	log.Infof("[AuthorManager.Get] Fetching author with ID: %d", idToGet)

	author, exists := m.authors[idToGet]
	if !exists {
		log.Errorf("[AuthorManager.Get] Author with ID %s not found", idToGet)
		return models.Author{}, fmt.Errorf("[AuthorManager.Get] Author with ID %d not found", idToGet)
	}

	log.Infof("[AuthorManager.Get] Successfully fetched author with ID: %d", idToGet)
	return author, nil
}

func (m *AuthorManager) Create(newAuthor models.Author) (models.Author, error) {
	log.Infof("[AuthorManager.Create] Creating new author")

	err := newAuthor.Validate()
	if err != nil {
		return models.Author{}, err
	}

	if newAuthor.ID == uuid.Nil {
		newAuthor.ID = uuid.New()
	}

	_, exists := m.authors[newAuthor.ID]
	if exists {
		log.Errorf("[AuthorManager.Create] Author with ID %s already exists", newAuthor.ID)
		return models.Author{}, fmt.Errorf("[AuthorManager.Create] Author with ID %d already exists", newAuthor.ID)
	}

	m.authors[newAuthor.ID] = newAuthor
	log.Infof("[AuthorManager.Create] Successfully created author with ID: %s", newAuthor.ID)
	return newAuthor, nil
}

func (m *AuthorManager) Update(updatedAuthor models.Author) (models.Author, error) {
	log.Infof("[AuthorManager.Update] Updating author with ID: %d", updatedAuthor.ID)

	_, exists := m.authors[updatedAuthor.ID]
	if !exists {
		log.Errorf("[AuthorManager.Update] Author with ID %s not found", updatedAuthor.ID)
		return models.Author{}, fmt.Errorf("[AuthorManager.Update] Author with ID %d not found", updatedAuthor.ID)
	}

	m.authors[updatedAuthor.ID] = updatedAuthor
	log.Infof("[AuthorManager.Update] Successfully updated author with ID: %s", updatedAuthor.ID)
	return updatedAuthor, nil
}

func (m *AuthorManager) Delete(idToDelete uuid.UUID) (models.Author, error) {
	log.Infof("[AuthorManager.Delete] Deleting author with ID: %d", idToDelete)

	deletedAuthor, exists := m.authors[idToDelete]
	if !exists {
		log.Errorf("[AuthorManager.Delete] Author with ID %s not found", idToDelete)
		return models.Author{}, fmt.Errorf("[AuthorManager.Delete] Author with ID %d not found", idToDelete)
	}

	delete(m.authors, idToDelete)
	log.Infof("[AuthorManager.Delete] Successfully deleted author with ID: %s", idToDelete)
	return deletedAuthor, nil
}
