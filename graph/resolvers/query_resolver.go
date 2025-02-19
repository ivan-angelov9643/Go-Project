package resolvers

import (
	"awesomeProject/graph/generated/graphql"
	"context"
)

type queryResolver struct {
	*RootResolver
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*graphql.Book, error) {
	return r.bookResolver.Books(ctx)
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id string) (*graphql.Book, error) {
	return r.bookResolver.Book(ctx, id)
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*graphql.Category, error) {
	return r.categoryResolver.Categories(ctx)
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*graphql.Category, error) {
	return r.categoryResolver.Category(ctx, id)
}
