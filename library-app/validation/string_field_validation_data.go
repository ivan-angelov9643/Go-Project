package validation

import (
	"fmt"
	"regexp"
)

type StringFieldValidationData struct {
	FieldName               string
	MinLength               *int
	MaxLength               *int
	CanBeEmpty              bool
	Regex                   *regexp.Regexp
	RegexFormatErrorMessage *string
}

func (data *StringFieldValidationData) GetFieldName() string {
	return data.FieldName
}

func (data *StringFieldValidationData) RegexErrorMessage() string {
	if data.RegexFormatErrorMessage == nil {
		return ""
	}
	return fmt.Sprintf(*data.RegexFormatErrorMessage, data.FieldName)
}

func (data *StringFieldValidationData) Validate(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string", data.FieldName)
	}

	if !data.CanBeEmpty && len(strValue) == 0 {
		return fmt.Errorf("%s cannot be empty", data.FieldName)
	}

	if data.MinLength != nil && len(strValue) < *data.MinLength {
		return fmt.Errorf("%s must be at least %d characters", data.FieldName, *data.MinLength)
	}

	if data.MaxLength != nil && len(strValue) > *data.MaxLength {
		return fmt.Errorf("%s must be no more than %d characters", data.FieldName, *data.MaxLength)
	}

	if data.Regex != nil && data.RegexFormatErrorMessage != nil && !data.Regex.MatchString(strValue) {
		return fmt.Errorf(data.RegexErrorMessage())
	}

	return nil
}
