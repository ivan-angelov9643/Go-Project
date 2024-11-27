package validation

import "fmt"

type FloatFieldValidationData struct {
	FieldName string
	MinValue  *float64
	MaxValue  *float64
}

func (data *FloatFieldValidationData) GetFieldName() string {
	return data.FieldName
}

func (data *FloatFieldValidationData) Validate(value interface{}) error {
	floatValue, ok := value.(float64)
	if !ok {
		return fmt.Errorf("%s must be a float", data.FieldName)
	}

	if data.MinValue != nil && floatValue < *data.MinValue {
		return fmt.Errorf("%s must be at least %.2f", data.FieldName, *data.MinValue)
	}

	if data.MaxValue != nil && floatValue > *data.MaxValue {
		return fmt.Errorf("%s must be no more than %.2f", data.FieldName, *data.MaxValue)
	}

	return nil
}
