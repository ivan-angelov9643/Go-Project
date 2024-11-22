package validation

import (
	"fmt"
	"regexp"
)

type FieldValidationData struct {
	FieldName               string
	MinLength               *int
	MaxLength               *int
	CanBeEmpty              bool
	Regex                   *regexp.Regexp
	RegexFormatErrorMessage *string
}

func (data *FieldValidationData) RegexErrorMessage() string {
	if data.RegexFormatErrorMessage == nil {
		return ""
	}
	return fmt.Sprintf(*data.RegexFormatErrorMessage, data.FieldName)
}

func (data *FieldValidationData) Validate(value string) error {
	if !data.CanBeEmpty && len(value) == 0 {
		return fmt.Errorf("%s cannot be empty", data.FieldName)
	}

	if data.MinLength != nil && len(value) < *data.MinLength {
		return fmt.Errorf("%s must be at least %d characters", data.FieldName, data.MinLength)
	}

	if data.MaxLength != nil && len(value) > *data.MaxLength {
		return fmt.Errorf("%s must be no more than %d characters", data.FieldName, data.MaxLength)
	}

	if data.Regex != nil && data.RegexFormatErrorMessage != nil && !data.Regex.MatchString(value) {
		return fmt.Errorf(data.RegexErrorMessage())
	}

	return nil
}
