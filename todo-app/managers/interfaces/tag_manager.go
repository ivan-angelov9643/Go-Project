package interfaces

import (
	"awesomeProject/todo-app/structs"
)

//go:generate mockery --name=TagManager --output=automock --with-expecter=true --outpkg=automock --case=underscore --disable-version-string
type TagManager interface {
	GetAll() []structs.Tag
	Get(string) (structs.Tag, error)
	Create(structs.Tag) (structs.Tag, error)
	Delete(string) (structs.Tag, error)
}
