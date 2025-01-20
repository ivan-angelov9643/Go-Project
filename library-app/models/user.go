package models

import (
	"awesomeProject/library-app/validation"
	log "github.com/sirupsen/logrus"
)

type User struct {
	BaseModel
	PreferredUsername string `json:"preferred_username"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	Email             string `json:"email"`
}

func (user *User) Validate() error {
	log.Info("[User.Validate] Validating user data")

	for _, validationData := range validation.UserValidation {
		var fieldValue interface{}

		switch validationData.GetFieldName() {
		case "Preferred UserName":
			fieldValue = user.PreferredUsername
		case "Given Name":
			fieldValue = user.GivenName
		case "Family Name":
			fieldValue = user.FamilyName
		case "Email":
			fieldValue = user.Email
		}

		err := validationData.Validate(fieldValue)
		if err != nil {
			log.Errorf("[User.Validate] Validation failed for field '%s' with value '%v': %v", validationData.GetFieldName(), fieldValue, err)
			return err
		}
	}

	log.Info("[User.Validate] Validation completed successfully")
	return nil
}
