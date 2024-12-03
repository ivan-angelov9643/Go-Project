package validation

import "fmt"

type IntFieldValidationData struct {
	FieldName string
	MinValue  *int
	MaxValue  *int
}

func (data *IntFieldValidationData) GetFieldName() string {
	return data.FieldName
}

func (data *IntFieldValidationData) Validate(value interface{}) error {
	intValue, ok := value.(int)
	if !ok {
		return fmt.Errorf("%s must be an int", data.FieldName)
	}

	if data.MinValue != nil && intValue < *data.MinValue {
		return fmt.Errorf("%s must be at least %d", data.FieldName, *data.MinValue)
	}

	if data.MaxValue != nil && intValue > *data.MaxValue {
		return fmt.Errorf("%s must be no more than %d", data.FieldName, *data.MaxValue)
	}

	return nil
}
