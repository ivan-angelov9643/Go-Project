package models

import (
	"awesomeProject/library-app/global/validation"
	log "github.com/sirupsen/logrus"
)

type Category struct {
	BaseModel
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"` // optional
}

func (category *Category) Validate() error {
	log.Info("[Category.Validate] Validating category data")

	for _, validationData := range validation.CategoryValidation {
		var fieldValue interface{}

		switch validationData.GetFieldName() {
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
			log.Errorf("[Category.Validate] Validation failed for field '%s' with value '%s': %v", validationData.GetFieldName(), fieldValue, err)
			return err
		}
	}

	log.Info("[Category.Validate] Validation completed successfully")
	return nil
}
