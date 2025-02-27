package managers

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ivan-angelov9643/go-project/library-app/db"
	"github.com/ivan-angelov9643/go-project/library-app/models"
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

func (m *AuthorManager) GetAll(scopes ...db.DBScope) ([]models.Author, error) {
	log.Info("[AuthorManager.GetAll] Fetching all authors")

	var allAuthors []models.Author
	err := db.ApplyScopes(m.db, scopes).Find(&allAuthors).Error
	if err != nil {
		log.Errorf("[AuthorManager.GetAll] Error fetching all authors: %v", err)
		return nil, db.NewDBError(db.InternalError, "[AuthorManager.GetAll] Error fetching all authors: %v", err)
	}

	log.Infof("[AuthorManager.GetAll] Successfully fetched all authors")
	return allAuthors, nil
}

func (m *AuthorManager) Get(idToGet uuid.UUID) (models.Author, error) {
	log.Infof("[AuthorManager.Get] Fetching author with ID: %s", idToGet)

	var author models.Author
	err := m.db.First(&author, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[AuthorManager.Get] Error fetching author with ID %s: %v", idToGet, err)
		return models.Author{}, db.NewDBError(db.InternalError, "[AuthorManager.Get] Error fetching author with ID %s: %v", idToGet, err)
	}

	log.Infof("[AuthorManager.Get] Successfully fetched author with ID: %s", idToGet)
	return author, nil
}

func (m *AuthorManager) Create(newAuthor models.Author) (models.Author, error) {
	log.Infof("[AuthorManager.Create] Creating new author")

	err := newAuthor.Validate()
	if err != nil {
		return models.Author{}, db.NewDBError(db.ValidationError, err.Error())
	}

	newAuthor.ID = uuid.New()

	err = m.db.Create(&newAuthor).Error
	if err != nil {
		log.Errorf("[AuthorManager.Create] Error creating new author with ID %s: %v", newAuthor.ID, err)
		return models.Author{}, db.NewDBError(db.InternalError, "[AuthorManager.Create] Error creating new author with ID %s: %v", newAuthor.ID, err)
	}

	log.Infof("[AuthorManager.Create] Successfully created author with ID: %s", newAuthor.ID)
	return newAuthor, nil
}

func (m *AuthorManager) Update(updatedAuthor models.Author) (models.Author, error) {
	log.Infof("[AuthorManager.Update] Updating author with ID: %s", updatedAuthor.ID)

	err := updatedAuthor.Validate()
	if err != nil {
		return models.Author{}, db.NewDBError(db.ValidationError, err.Error())
	}

	var author models.Author
	err = m.db.First(&author, "id = ?", updatedAuthor.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[AuthorManager.Update] Author with ID %s does not exist", updatedAuthor.ID)
			return models.Author{}, db.NewDBError(db.NotFoundError, "[AuthorManager.Update] Author with ID %s does not exist", updatedAuthor.ID)
		}
		log.Errorf("[AuthorManager.Update] Error fetching author with ID %s: %v", updatedAuthor.ID, err)
		return models.Author{}, db.NewDBError(db.InternalError, "[AuthorManager.Update] Error fetching author with ID %s: %v", updatedAuthor.ID, err)
	}

	err = m.db.Model(&author).Updates(updatedAuthor).Error
	if err != nil {
		log.Errorf("[AuthorManager.Update] Error updating author with ID %s: %v", updatedAuthor.ID, err)
		return models.Author{}, db.NewDBError(db.InternalError, "[AuthorManager.Update] Error updating author with ID %s: %v", updatedAuthor.ID, err)
	}

	log.Infof("[AuthorManager.Update] Successfully updated author with ID: %s", updatedAuthor.ID)
	return updatedAuthor, nil
}

func (m *AuthorManager) Delete(idToDelete uuid.UUID) (models.Author, error) {
	log.Infof("[AuthorManager.Delete] Deleting author with ID: %s", idToDelete)

	var author models.Author
	err := m.db.First(&author, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[AuthorManager.Delete] Author with ID %s does not exist", idToDelete)
			return models.Author{}, db.NewDBError(db.NotFoundError, "[AuthorManager.Delete] Author with ID %s does not exist", idToDelete)
		}
		log.Errorf("[AuthorManager.Delete] Error fetching author with ID %s: %v", idToDelete, err)
		return models.Author{}, db.NewDBError(db.InternalError, "[AuthorManager.Delete] Error fetching author with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&author).Error
	if err != nil {
		log.Errorf("[AuthorManager.Delete] Error deleting author with ID %s: %v", idToDelete, err)
		return models.Author{}, db.NewDBError(db.InternalError, "[AuthorManager.Delete] Error deleting author with ID %s: %v", idToDelete, err)
	}

	log.Infof("[AuthorManager.Delete] Successfully deleted author with ID: %s", idToDelete)
	return author, nil
}

func (m *AuthorManager) Count(scopes ...db.DBScope) (int64, error) {
	log.Infof("[AuthorManager.Count] Counting authors in the database")

	var count int64
	err := db.ApplyScopes(m.db, scopes).Model(&models.Author{}).Count(&count).Error
	if err != nil {
		log.Errorf("[AuthorManager.Count] Error counting authors: %v", err)
		return 0, db.NewDBError(db.InternalError, "[AuthorManager.Count] Error counting authors: %v", err)
	}

	log.Infof("[AuthorManager.Count] Successfully counted authors: %d", count)
	return count, nil
}
