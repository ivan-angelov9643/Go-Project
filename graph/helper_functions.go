package graph

import (
	"awesomeProject/graph/model"
	"awesomeProject/library-app/models"
	"fmt"
	"github.com/google/uuid"
)

func GORMBookModel(title string, year int32, authorID string, categoryID string, totalCopies int32, language string) (*models.Book, error) {
	authorUUID, err := uuid.Parse(authorID)
	if err != nil {
		return nil, fmt.Errorf("invalid author ID: %w", err)
	}

	categoryUUID, err := uuid.Parse(categoryID)
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}

	return &models.Book{
		Title:       title,
		Year:        int(year),
		AuthorID:    authorUUID,
		CategoryID:  categoryUUID,
		TotalCopies: int(totalCopies),
		Language:    language,
	}, nil
}

func ToGraphQLBookModel(book models.Book) model.Book {
	return model.Book{
		ID:              book.ID.String(),
		Title:           book.Title,
		Year:            int32(book.Year),
		Author:          book.AuthorID.String(),
		Category:        &model.Category{ID: book.CategoryID.String(), Name: book.CategoryName},
		TotalCopies:     int32(book.TotalCopies),
		AvailableCopies: int32(book.AvailableCopies),
		Language:        book.Language,
	}
}

func GORMCategoryModel(name string, description *string) *models.Category {
	return &models.Category{
		Name:        name,
		Description: description,
	}
}

func ToGraphQLCategoryModel(category models.Category) model.Category {
	return model.Category{
		ID:          category.ID.String(),
		Name:        category.Name,
		Description: category.Description,
	}
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefInt(i *int32) int {
	if i == nil {
		return 0
	}
	return int(*i)
}
