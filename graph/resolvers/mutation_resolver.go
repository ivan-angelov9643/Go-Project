package resolvers

import (
	"awesomeProject/graph/generated/graphql"
	"context"
)

type mutationResolver struct {
	*RootResolver
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, title string, year int32, authorID string, categoryID string, totalCopies int32, language string) (*graphql.Book, error) {
	return r.bookResolver.CreateBook(ctx, title, year, authorID, categoryID, totalCopies, language)
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, title *string, year *int32, totalCopies *int32, language *string) (*graphql.Book, error) {
	return r.bookResolver.UpdateBook(ctx, id, title, year, totalCopies, language)
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (bool, error) {
	return r.bookResolver.DeleteBook(ctx, id)
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, name string, description *string) (*graphql.Category, error) {
	return r.categoryResolver.CreateCategory(ctx, name, description)
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, name *string, description *string) (*graphql.Category, error) {
	return r.categoryResolver.UpdateCategory(ctx, id, name, description)
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	return r.categoryResolver.DeleteCategory(ctx, id)
}
