package implementations

import (
	"awesomeProject/todo-app/models"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BookManager struct {
	db    *gorm.DB
	books map[uuid.UUID]models.Book
}

func NewBookManager(db *gorm.DB) *BookManager {
	log.Info("[NewBookManager] Initializing BookManager")

	return &BookManager{db, make(map[uuid.UUID]models.Book)}
}

func (m *BookManager) GetAll() []models.Book {
	log.Info("[BookManager.GetAll] Fetching all books")

	allBooks := make([]models.Book, 0, len(m.books))
	for _, book := range m.books {
		allBooks = append(allBooks, book)
	}

	log.Infof("[BookManager.GetAll] Successfully fetched all books")
	return allBooks
}

func (m *BookManager) Get(idToGet uuid.UUID) (models.Book, error) {
	log.Infof("[BookManager.Get] Fetching book with ID: %s", idToGet)

	book, exists := m.books[idToGet]
	if !exists {
		log.Errorf("[BookManager.Get] Book with ID %s not found", idToGet)
		return models.Book{}, fmt.Errorf("[BookManager.Get] Book with ID %s not found", idToGet)
	}

	log.Infof("[BookManager.Get] Successfully fetched book with ID: %s", idToGet)
	return book, nil
}

func (m *BookManager) Create(newBook models.Book) (models.Book, error) {
	log.Infof("[BookManager.Create] Creating new book")

	if newBook.ID == uuid.Nil {
		newBook.ID = uuid.New()
	}

	_, exists := m.books[newBook.ID]
	if exists {
		log.Errorf("[BookManager.Create] Book with ID %s already exists", newBook.ID)
		return models.Book{}, fmt.Errorf("[BookManager.Create] Book with ID %s already exists", newBook.ID)
	}

	m.books[newBook.ID] = newBook
	log.Infof("[BookManager.Create] Successfully created book with ID: %s", newBook.ID)
	return newBook, nil
}

func (m *BookManager) Update(updatedBook models.Book) (models.Book, error) {
	log.Infof("[BookManager.Update] Updating book with ID: %s", updatedBook.ID)

	_, exists := m.books[updatedBook.ID]
	if !exists {
		log.Errorf("[BookManager.Update] Book with ID %s not found", updatedBook.ID)
		return models.Book{}, fmt.Errorf("[BookManager.Update] Book with ID %s not found", updatedBook.ID)
	}

	m.books[updatedBook.ID] = updatedBook
	log.Infof("[BookManager.Update] Successfully updated book with ID: %s", updatedBook.ID)
	return updatedBook, nil
}

func (m *BookManager) Delete(idToDelete uuid.UUID) (models.Book, error) {
	log.Infof("[BookManager.Delete] Deleting book with ID: %s", idToDelete)

	deletedBook, exists := m.books[idToDelete]
	if !exists {
		log.Errorf("[BookManager.Delete] Book with ID %s not found", idToDelete)
		return models.Book{}, fmt.Errorf("[BookManager.Delete] Book with ID %s not found", idToDelete)
	}

	delete(m.books, idToDelete)
	log.Infof("[BookManager.Delete] Successfully deleted book with ID: %s", idToDelete)
	return deletedBook, nil
}
