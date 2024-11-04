package interfaces

import (
	"awesomeProject/todo-app/structs"
	"github.com/google/uuid"
)

//go:generate mockery --name=ListManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type ListManager interface {
	GetAll() []structs.List
	Get(uuid.UUID) (structs.List, error)
	Create(structs.List) (structs.List, error)
	Update(structs.List) (structs.List, error)
	Delete(uuid.UUID) (structs.List, error)
}
