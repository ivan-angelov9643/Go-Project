package implementations

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthorManager struct {
	db *gorm.DB
}

func NewAuthorManager(db *gorm.DB) *AuthorManager {
	log.Info("[NewAuthorManager] Initializing AuthorManager")

	return &AuthorManager{db}
}

func (m *AuthorManager) GetAll() ([]models.Author, *db_error.DBError) {
	log.Info("[AuthorManager.GetAll] Fetching all authors")

	var allAuthors []models.Author
	err := m.db.Find(&allAuthors).Error
	if err != nil {
		log.Errorf("[AuthorManager.GetAll] Error fetching all authors: %v", err)
		return nil, db_error.NewDBError(db_error.InternalError, "[AuthorManager.GetAll] Error fetching all authors: %v", err)
	}

	log.Infof("[AuthorManager.GetAll] Successfully fetched all authors")
	return allAuthors, nil
}

func (m *AuthorManager) Get(idToGet uuid.UUID) (models.Author, *db_error.DBError) {
	log.Infof("[AuthorManager.Get] Fetching author with ID: %s", idToGet)

	var author models.Author
	err := m.db.First(&author, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[AuthorManager.Get] Error fetching author with ID %s: %v", idToGet, err)
		return models.Author{}, db_error.NewDBError(db_error.InternalError, "[AuthorManager.Get] Error fetching author with ID %s: %v", idToGet, err)
	}

	log.Infof("[AuthorManager.Get] Successfully fetched author with ID: %s", idToGet)
	return author, nil
}

func (m *AuthorManager) Create(newAuthor models.Author) (models.Author, *db_error.DBError) {
	log.Infof("[AuthorManager.Create] Creating new author")

	err := newAuthor.Validate()
	if err != nil {
		return models.Author{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	newAuthor.ID = uuid.New()

	err = m.db.Create(&newAuthor).Error
	if err != nil {
		log.Errorf("[AuthorManager.Create] Error creating new author with ID %s: %v", newAuthor.ID, err)
		return models.Author{}, db_error.NewDBError(db_error.InternalError, "[AuthorManager.Create] Error creating new author with ID %s: %v", newAuthor.ID, err)
	}

	log.Infof("[AuthorManager.Create] Successfully created author with ID: %s", newAuthor.ID)
	return newAuthor, nil
}

func (m *AuthorManager) Update(updatedAuthor models.Author) (models.Author, *db_error.DBError) {
	log.Infof("[AuthorManager.Update] Updating author with ID: %s", updatedAuthor.ID)

	err := updatedAuthor.Validate()
	if err != nil {
		return models.Author{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	var author models.Author
	err = m.db.First(&author, "id = ?", updatedAuthor.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[AuthorManager.Update] Author with ID %s does not exist", updatedAuthor.ID)
			return models.Author{}, db_error.NewDBError(db_error.NotFoundError, "[AuthorManager.Update] Author with ID %s does not exist", updatedAuthor.ID)
		}
		log.Errorf("[AuthorManager.Update] Error fetching author with ID %s: %v", updatedAuthor.ID, err)
		return models.Author{}, db_error.NewDBError(db_error.InternalError, "[AuthorManager.Update] Error fetching author with ID %s: %v", updatedAuthor.ID, err)
	}

	err = m.db.Model(&author).Updates(updatedAuthor).Error
	if err != nil {
		log.Errorf("[AuthorManager.Update] Error updating author with ID %s: %v", updatedAuthor.ID, err)
		return models.Author{}, db_error.NewDBError(db_error.InternalError, "[AuthorManager.Update] Error updating author with ID %s: %v", updatedAuthor.ID, err)
	}

	log.Infof("[AuthorManager.Update] Successfully updated author with ID: %s", updatedAuthor.ID)
	return updatedAuthor, nil
}

func (m *AuthorManager) Delete(idToDelete uuid.UUID) (models.Author, *db_error.DBError) {
	log.Infof("[AuthorManager.Delete] Deleting author with ID: %s", idToDelete)

	var author models.Author
	err := m.db.First(&author, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[AuthorManager.Delete] Author with ID %s does not exist", idToDelete)
			return models.Author{}, db_error.NewDBError(db_error.NotFoundError, "[AuthorManager.Delete] Author with ID %s does not exist", idToDelete)
		}
		log.Errorf("[AuthorManager.Delete] Error fetching author with ID %s: %v", idToDelete, err)
		return models.Author{}, db_error.NewDBError(db_error.InternalError, "[AuthorManager.Delete] Error fetching author with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&author).Error
	if err != nil {
		log.Errorf("[AuthorManager.Delete] Error deleting author with ID %s: %v", idToDelete, err)
		return models.Author{}, db_error.NewDBError(db_error.InternalError, "[AuthorManager.Delete] Error deleting author with ID %s: %v", idToDelete, err)
	}

	log.Infof("[AuthorManager.Delete] Successfully deleted author with ID: %s", idToDelete)
	return author, nil
}
