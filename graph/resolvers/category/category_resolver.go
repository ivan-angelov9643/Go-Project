package category

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
	categoryManager managers.CategoryManager
}

func NewResolver(manager managers.CategoryManager) *Resolver {
	return &Resolver{
		categoryManager: manager,
	}
}

// CreateCategory is the resolver for the createCategory field.
func (r *Resolver) CreateCategory(ctx context.Context, name string, description *string) (*graphql.Category, error) {
	category := graph.GORMCategoryModel(name, description)

	newCategory, err := r.categoryManager.Create(*category)
	if err != nil {
		return nil, err
	}

	graphQLCategory := graph.ToGraphQLCategoryModel(newCategory)
	return &graphQLCategory, nil
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *Resolver) UpdateCategory(ctx context.Context, id string, name *string, description *string) (*graphql.Category, error) {
	categoryID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}

	updatedCategory := models.Category{
		BaseModel:   models.BaseModel{ID: categoryID},
		Name:        graph.DerefString(name),
		Description: description,
	}

	category, err := r.categoryManager.Update(updatedCategory)
	if err != nil {
		return nil, err
	}

	graphQLCategory := graph.ToGraphQLCategoryModel(category)
	return &graphQLCategory, nil
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *Resolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	categoryID, err := uuid.Parse(id)
	if err != nil {
		return false, fmt.Errorf("invalid category ID: %w", err)
	}

	_, err = r.categoryManager.Delete(categoryID)
	return err == nil, err
}

// Categories is the resolver for the categories field.
func (r *Resolver) Categories(ctx context.Context) ([]*graphql.Category, error) {
	categories, err := r.categoryManager.GetAll()
	if err != nil {
		return nil, err
	}

	var graphQLCategories []*graphql.Category
	for _, category := range categories {
		c := graph.ToGraphQLCategoryModel(category)
		graphQLCategories = append(graphQLCategories, &c)
	}

	return graphQLCategories, nil
}

// Category is the resolver for the category field.
func (r *Resolver) Category(ctx context.Context, id string) (*graphql.Category, error) {
	categoryID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}

	category, err := r.categoryManager.Get(categoryID)
	if err != nil {
		return nil, err
	}

	graphQLCategory := graph.ToGraphQLCategoryModel(category)
	return &graphQLCategory, nil
}
