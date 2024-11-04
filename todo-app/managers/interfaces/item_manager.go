package interfaces

import (
	"awesomeProject/todo-app/structs"
	"github.com/google/uuid"
)

//go:generate mockery --name=ItemManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ItemManager interface {
	GetAll() []structs.Item
	Get(uuid.UUID) (structs.Item, error)
	Create(structs.Item) (structs.Item, error)
	Update(structs.Item) (structs.Item, error)
	Delete(uuid.UUID) (structs.Item, error)
}
