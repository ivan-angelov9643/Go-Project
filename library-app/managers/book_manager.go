package managers

import (
	"awesomeProject/library-app/db"
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

func (m *BookManager) GetAll(scopes ...db.DBScope) ([]models.Book, error) {
	log.Info("[BookManager.GetAll] Fetching all books")

	var allBooks []models.Book
	err := db.ApplyScopes(m.db, scopes).Table("books").
		Select(`books.*, 
            CONCAT(authors.first_name, ' ', authors.last_name) AS author_name, 
            categories.name AS category_name,
            (books.total_copies - 
                (SELECT COUNT(*) FROM reservations WHERE reservations.book_id = books.id) - 
                (SELECT COUNT(*) FROM loans WHERE loans.book_id = books.id AND loans.status = 'active')
            ) AS available_copies`).
		Joins("LEFT JOIN authors ON authors.id = books.author_id").
		Joins("LEFT JOIN categories ON categories.id = books.category_id").
		Find(&allBooks).Error
	if err != nil {
		log.Errorf("[BookManager.GetAll] Error fetching all books: %v", err)
		return nil, db.NewDBError(db.InternalError, "[BookManager.GetAll] Error fetching all books: %v", err)
	}

	log.Infof("[BookManager.GetAll] Successfully fetched all books")
	return allBooks, nil
}

func (m *BookManager) Get(idToGet uuid.UUID) (models.Book, error) {
	log.Infof("[BookManager.Get] Fetching book with ID: %s", idToGet)

	var book models.Book
	err := m.db.Table("books").
		Select(`books.*, 
            CONCAT(authors.first_name, ' ', authors.last_name) AS author_name, 
            categories.name AS category_name,
            (books.total_copies - 
                (SELECT COUNT(*) FROM reservations WHERE reservations.book_id = books.id) - 
                (SELECT COUNT(*) FROM loans WHERE loans.book_id = books.id AND loans.status = 'active')
            ) AS available_copies`).
		Joins("LEFT JOIN authors ON authors.id = books.author_id").
		Joins("LEFT JOIN categories ON categories.id = books.category_id").
		Where("books.id = ?", idToGet).
		First(&book).Error
	if err != nil {
		log.Errorf("[BookManager.Get] Error fetching book with ID %s: %v", idToGet, err)
		return models.Book{}, db.NewDBError(db.InternalError, "[BookManager.Get] Error fetching book with ID %s: %v", idToGet, err)
	}

	log.Infof("[BookManager.Get] Successfully fetched book with ID: %s", idToGet)
	return book, nil
}

func (m *BookManager) Create(newBook models.Book) (models.Book, error) {
	log.Infof("[BookManager.Create] Creating new book")

	err := newBook.Validate()
	if err != nil {
		return models.Book{}, db.NewDBError(db.ValidationError, err.Error())
	}

	newBook.ID = uuid.New()

	err = m.db.Create(&newBook).Error
	if err != nil {
		log.Errorf("[BookManager.Create] Error creating new book with ID %s: %v", newBook.ID, err)
		return models.Book{}, db.NewDBError(db.InternalError, "[BookManager.Create] Error creating new book with ID %s: %v", newBook.ID, err)
	}

	log.Infof("[BookManager.Create] Successfully created book with ID: %s", newBook.ID)
	return newBook, nil
}

func (m *BookManager) Update(updatedBook models.Book) (models.Book, error) {
	log.Infof("[BookManager.Update] Updating book with ID: %s", updatedBook.ID)

	err := updatedBook.Validate()
	if err != nil {
		return models.Book{}, db.NewDBError(db.ValidationError, err.Error())
	}

	var book models.Book
	err = m.db.First(&book, "id = ?", updatedBook.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[BookManager.Update] Book with ID %s does not exist", updatedBook.ID)
			return models.Book{}, db.NewDBError(db.NotFoundError, "[BookManager.Update] Book with ID %s does not exist", updatedBook.ID)
		}
		log.Errorf("[BookManager.Update] Error fetching book with ID %s: %v", updatedBook.ID, err)
		return models.Book{}, db.NewDBError(db.InternalError, "[BookManager.Update] Error fetching book with ID %s: %v", updatedBook.ID, err)
	}

	err = m.db.Model(&book).Updates(updatedBook).Error
	if err != nil {
		log.Errorf("[BookManager.Update] Error updating book with ID %s: %v", updatedBook.ID, err)
		return models.Book{}, db.NewDBError(db.InternalError, "[BookManager.Update] Error updating book with ID %s: %v", updatedBook.ID, err)
	}

	log.Infof("[BookManager.Update] Successfully updated book with ID: %s", updatedBook.ID)
	return updatedBook, nil
}

func (m *BookManager) Delete(idToDelete uuid.UUID) (models.Book, error) {
	log.Infof("[BookManager.Delete] Deleting book with ID: %s", idToDelete)

	var book models.Book
	err := m.db.First(&book, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[BookManager.Delete] Book with ID %s does not exist", idToDelete)
			return models.Book{}, db.NewDBError(db.NotFoundError, "[BookManager.Delete] Book with ID %s does not exist", idToDelete)
		}
		log.Errorf("[BookManager.Delete] Error fetching book with ID %s: %v", idToDelete, err)
		return models.Book{}, db.NewDBError(db.InternalError, "[BookManager.Delete] Error fetching book with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&book).Error
	if err != nil {
		log.Errorf("[BookManager.Delete] Error deleting book with ID %s: %v", idToDelete, err)
		return models.Book{}, db.NewDBError(db.InternalError, "[BookManager.Delete] Error deleting book with ID %s: %v", idToDelete, err)
	}

	log.Infof("[BookManager.Delete] Successfully deleted book with ID: %s", idToDelete)
	return book, nil
}

func (m *BookManager) Count(scopes ...db.DBScope) (int64, error) {
	log.Infof("[BookManager.Count] Counting books in the database")

	var count int64
	err := db.ApplyScopes(m.db, scopes).Model(&models.Book{}).Count(&count).Error
	if err != nil {
		log.Errorf("[BookManager.Count] Error counting books: %v", err)
		return 0, db.NewDBError(db.InternalError, "[BookManager.Count] Error counting books: %v", err)
	}

	log.Infof("[BookManager.Count] Successfully counted books: %d", count)
	return count, nil
}
