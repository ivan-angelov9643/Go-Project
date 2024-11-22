package validation

import (
	"awesomeProject/todo-app/global"
	"regexp"
)

var (
	NameRegex        = "^[A-Z]{1}[a-z]+$"
	NameRegexMessage = "%s must only contain uppercase letters followed by lowercase English letters only."
	AuthorValidation = []FieldValidationData{
		{
			FieldName:               "First Name",
			MinLength:               global.IntPtr(2),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(NameRegex),
			RegexFormatErrorMessage: &NameRegexMessage,
		},
		{
			FieldName:               "Last Name",
			MinLength:               global.IntPtr(2),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(NameRegex),
			RegexFormatErrorMessage: &NameRegexMessage,
		},
		{
			FieldName:               "Nationality",
			MinLength:               global.IntPtr(2),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(NameRegex),
			RegexFormatErrorMessage: &NameRegexMessage,
		},
		{
			FieldName:               "Bio",
			MinLength:               nil,
			MaxLength:               global.IntPtr(5000),
			CanBeEmpty:              true,
			Regex:                   nil,
			RegexFormatErrorMessage: nil,
		},
		{
			FieldName:               "Website",
			MinLength:               nil,
			MaxLength:               global.IntPtr(500),
			CanBeEmpty:              true,
			Regex:                   nil,
			RegexFormatErrorMessage: nil,
		},
	}
	CategoryValidation = []FieldValidationData{
		{
			FieldName:               "Name",
			MinLength:               global.IntPtr(2),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   nil,
			RegexFormatErrorMessage: nil,
		},
		{
			FieldName:               "Bio",
			MinLength:               nil,
			MaxLength:               global.IntPtr(5000),
			CanBeEmpty:              true,
			Regex:                   nil,
			RegexFormatErrorMessage: nil,
		},
	}
)
