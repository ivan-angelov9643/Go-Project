package implementations

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BookManager struct {
	db *gorm.DB
}

func NewBookManager(db *gorm.DB) *BookManager {
	log.Info("[NewBookManager] Initializing BookManager")

	return &BookManager{db}
}

func (m *BookManager) GetAll() ([]models.Book, *db_error.DBError) {
	log.Info("[BookManager.GetAll] Fetching all books")

	var allBooks []models.Book
	err := m.db.Find(&allBooks).Error
	if err != nil {
		log.Errorf("[BookManager.GetAll] Error fetching all books: %v", err)
		return nil, db_error.NewDBError(db_error.InternalError, "[BookManager.GetAll] Error fetching all books: %v", err)
	}

	log.Infof("[BookManager.GetAll] Successfully fetched all books")
	return allBooks, nil
}

func (m *BookManager) Get(idToGet uuid.UUID) (models.Book, *db_error.DBError) {
	log.Infof("[BookManager.Get] Fetching book with ID: %s", idToGet)

	var book models.Book
	err := m.db.First(&book, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[BookManager.Get] Error fetching book with ID %s: %v", idToGet, err)
		return models.Book{}, db_error.NewDBError(db_error.InternalError, "[BookManager.Get] Error fetching book with ID %s: %v", idToGet, err)
	}

	log.Infof("[BookManager.Get] Successfully fetched book with ID: %s", idToGet)
	return book, nil
}

func (m *BookManager) Create(newBook models.Book) (models.Book, *db_error.DBError) {
	log.Infof("[BookManager.Create] Creating new book")

	err := newBook.Validate()
	if err != nil {
		return models.Book{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	newBook.ID = uuid.New()

	err = m.db.Create(&newBook).Error
	if err != nil {
		log.Errorf("[BookManager.Create] Error creating new book with ID %s: %v", newBook.ID, err)
		return models.Book{}, db_error.NewDBError(db_error.InternalError, "[BookManager.Create] Error creating new book with ID %s: %v", newBook.ID, err)
	}

	log.Infof("[BookManager.Create] Successfully created book with ID: %s", newBook.ID)
	return newBook, nil
}

func (m *BookManager) Update(updatedBook models.Book) (models.Book, *db_error.DBError) {
	log.Infof("[BookManager.Update] Updating book with ID: %s", updatedBook.ID)

	err := updatedBook.Validate()
	if err != nil {
		return models.Book{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	var book models.Book
	err = m.db.First(&book, "id = ?", updatedBook.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[BookManager.Update] Book with ID %s does not exist", updatedBook.ID)
			return models.Book{}, db_error.NewDBError(db_error.NotFoundError, "[BookManager.Update] Book with ID %s does not exist", updatedBook.ID)
		}
		log.Errorf("[BookManager.Update] Error fetching book with ID %s: %v", updatedBook.ID, err)
		return models.Book{}, db_error.NewDBError(db_error.InternalError, "[BookManager.Update] Error fetching book with ID %s: %v", updatedBook.ID, err)
	}

	err = m.db.Model(&book).Updates(updatedBook).Error
	if err != nil {
		log.Errorf("[BookManager.Update] Error updating book with ID %s: %v", updatedBook.ID, err)
		return models.Book{}, db_error.NewDBError(db_error.InternalError, "[BookManager.Update] Error updating book with ID %s: %v", updatedBook.ID, err)
	}

	log.Infof("[BookManager.Update] Successfully updated book with ID: %s", updatedBook.ID)
	return updatedBook, nil
}

func (m *BookManager) Delete(idToDelete uuid.UUID) (models.Book, *db_error.DBError) {
	log.Infof("[BookManager.Delete] Deleting book with ID: %s", idToDelete)

	var book models.Book
	err := m.db.First(&book, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[BookManager.Delete] Book with ID %s does not exist", idToDelete)
			return models.Book{}, db_error.NewDBError(db_error.NotFoundError, "[BookManager.Delete] Book with ID %s does not exist", idToDelete)
		}
		log.Errorf("[BookManager.Delete] Error fetching book with ID %s: %v", idToDelete, err)
		return models.Book{}, db_error.NewDBError(db_error.InternalError, "[BookManager.Delete] Error fetching book with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&book).Error
	if err != nil {
		log.Errorf("[BookManager.Delete] Error deleting book with ID %s: %v", idToDelete, err)
		return models.Book{}, db_error.NewDBError(db_error.InternalError, "[BookManager.Delete] Error deleting book with ID %s: %v", idToDelete, err)
	}

	log.Infof("[BookManager.Delete] Successfully deleted book with ID: %s", idToDelete)
	return book, nil
}
