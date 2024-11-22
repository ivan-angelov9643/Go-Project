package interfaces

import (
	"awesomeProject/todo-app/models"
	"github.com/google/uuid"
)

//go:generate mockery --name=BookManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type BookManager interface {
	GetAll() []models.Book
	Get(uuid uuid.UUID) (models.Book, error)
	Create(models.Book) (models.Book, error)
	Update(models.Book) (models.Book, error)
	Delete(uuid uuid.UUID) (models.Book, error)
}
