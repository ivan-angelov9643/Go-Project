package resolvers

import (
	graph "awesomeProject/graph/generated"
	"awesomeProject/graph/generated/graphql"
	"awesomeProject/graph/resolvers/book"
	"awesomeProject/graph/resolvers/category"
	"awesomeProject/library-app/managers"
	"context"
	"gorm.io/gorm"
)

type RootResolver struct {
	bookResolver     *book.Resolver
	categoryResolver *category.Resolver
}

func NewRootResolver(db *gorm.DB) *RootResolver {
	bookManager := *managers.NewBookManager(db)
	categoryManager := *managers.NewCategoryManager(db)

	return &RootResolver{
		bookResolver:     book.NewResolver(bookManager),
		categoryResolver: category.NewResolver(categoryManager),
	}
}

type bookResolver struct{ *RootResolver }

func (r *RootResolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

func (r *RootResolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

// Book returns BookResolver implementation.
func (r *RootResolver) Book() graph.BookResolver { return &bookResolver{r} }

// Category is the resolver for the category field.
func (r *bookResolver) Category(ctx context.Context, obj *graphql.Book) (*graphql.Category, error) {
	return r.categoryResolver.Category(ctx, obj.Category.ID)
}
