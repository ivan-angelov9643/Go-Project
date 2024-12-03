package validation

type FieldValidationData interface {
	Validate(value interface{}) error
	GetFieldName() string
}
