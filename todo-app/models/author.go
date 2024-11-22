package models

import (
	"awesomeProject/todo-app/global/validation"
	"time"
)

type Author struct {
	BaseModel
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Nationality string     `json:"nationality"`
	BirthDate   time.Time  `json:"birth_date"`
	DeathDate   *time.Time `json:"death_date,omitempty"` // nullable
	Bio         *string    `json:"bio,omitempty"`        // optional
	Website     *string    `json:"website,omitempty"`    // nullable
}

func (author *Author) Validate() error {
	for _, validationData := range validation.AuthorValidation {
		var fieldValue string
		switch validationData.FieldName {
		case "First Name":
			fieldValue = author.FirstName
		case "Last Name":
			fieldValue = author.LastName
		case "Nationality":
			fieldValue = author.Nationality
		case "Bio":
			if author.Bio == nil {
				continue
			}
			fieldValue = *author.Bio
		case "Website":
			if author.Website == nil {
				continue
			}
			fieldValue = *author.Website
		}

		err := validationData.Validate(fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}
