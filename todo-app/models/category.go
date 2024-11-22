package models

import "awesomeProject/todo-app/global/validation"

type Category struct {
	BaseModel
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"` // optional
}

func (category *Category) Validate() error {
	for _, validationData := range validation.CategoryValidation {
		var fieldValue string
		switch validationData.FieldName {
		case "Name":
			fieldValue = category.Name
		case "Description":
			if category.Description == nil {
				continue
			}
			fieldValue = *category.Description
		}

		err := validationData.Validate(fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}
