package graph

import "github.com/ivan-angelov9643/go-project/library-app/managers"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BookManager     managers.BookManager
	CategoryManager managers.CategoryManager
}
