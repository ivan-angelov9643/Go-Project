package book

import (
	"awesomeProject/graph"
	"awesomeProject/graph/generated/graphql"
	"awesomeProject/library-app/managers"
	"awesomeProject/library-app/models"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Resolver struct {
	bookManager managers.BookManager
}

func NewResolver(manager managers.BookManager) *Resolver {
	return &Resolver{
		bookManager: manager,
	}
}

// CreateBook is the resolver for the createBook field.
func (r *Resolver) CreateBook(ctx context.Context, title string, year int32, authorID string, categoryID string, totalCopies int32, language string) (*graphql.Book, error) {
	book, err := graph.GORMBookModel(title, year, authorID, categoryID, totalCopies, language)
	if err != nil {
		return nil, err
	}

	newBook, err := r.bookManager.Create(*book)
	if err != nil {
		return nil, err
	}

	graphQLBook := graph.ToGraphQLBookModel(newBook)
	return &graphQLBook, nil
}

// UpdateBook is the resolver for the updateBook field.
func (r *Resolver) UpdateBook(ctx context.Context, id string, title *string, year *int32, totalCopies *int32, language *string) (*graphql.Book, error) {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %w", err)
	}

	updatedBook := models.Book{
		BaseModel:   models.BaseModel{ID: bookID},
		Title:       graph.DerefString(title),
		Year:        graph.DerefInt(year),
		TotalCopies: graph.DerefInt(totalCopies),
		Language:    graph.DerefString(language),
	}

	book, err := r.bookManager.Update(updatedBook)
	if err != nil {
		return nil, err
	}

	graphQLBook := graph.ToGraphQLBookModel(book)
	return &graphQLBook, nil
}

// DeleteBook is the resolver for the deleteBook field.
func (r *Resolver) DeleteBook(ctx context.Context, id string) (bool, error) {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return false, fmt.Errorf("invalid book ID: %w", err)
	}

	_, err = r.bookManager.Delete(bookID)
	return err == nil, err
}

// Books is the resolver for the books field.
func (r *Resolver) Books(ctx context.Context) ([]*graphql.Book, error) {
	books, err := r.bookManager.GetAll()
	if err != nil {
		return nil, err
	}

	var graphQLBooks []*graphql.Book
	for _, book := range books {
		b := graph.ToGraphQLBookModel(book)
		graphQLBooks = append(graphQLBooks, &b)
	}

	return graphQLBooks, nil
}

// Book is the resolver for the book field.
func (r *Resolver) Book(ctx context.Context, id string) (*graphql.Book, error) {
	bookID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %w", err)
	}

	book, err := r.bookManager.Get(bookID)
	if err != nil {
		return nil, err
	}

	graphQLBook := graph.ToGraphQLBookModel(book)
	return &graphQLBook, nil
}
