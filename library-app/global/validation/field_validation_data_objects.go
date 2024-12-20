package validation

import (
	"awesomeProject/library-app/global"
	"regexp"
	"time"
)

var (
	NameRegex        = "^[A-Z]{1}[a-z]+$"
	NameRegexMessage = "%s must start with uppercase letter followed by lowercase English letters only."
	UserValidation   = []FieldValidationData{
		&StringFieldValidationData{
			FieldName:               "Preferred UserName",
			MinLength:               global.IntPtr(1),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile("^[A-Za-z0-9_]+$"),
			RegexFormatErrorMessage: global.StrPtr("Preferred user name must be alphanumeric and can include underscores."),
		},
		&StringFieldValidationData{
			FieldName:               "Given Name",
			MinLength:               global.IntPtr(1),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(NameRegex),
			RegexFormatErrorMessage: &NameRegexMessage,
		},
		&StringFieldValidationData{
			FieldName:               "Family Name",
			MinLength:               global.IntPtr(1),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(NameRegex),
			RegexFormatErrorMessage: &NameRegexMessage,
		},
		&StringFieldValidationData{
			FieldName:               "Email",
			MinLength:               global.IntPtr(5),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
			RegexFormatErrorMessage: global.StrPtr("Email must be a valid email address."),
		},
	}
	AuthorValidation = []FieldValidationData{
		&StringFieldValidationData{
			FieldName:               "First Name",
			MinLength:               global.IntPtr(2),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(NameRegex),
			RegexFormatErrorMessage: &NameRegexMessage,
		},
		&StringFieldValidationData{
			FieldName:               "Last Name",
			MinLength:               global.IntPtr(2),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(NameRegex),
			RegexFormatErrorMessage: &NameRegexMessage,
		},
		&StringFieldValidationData{
			FieldName:               "Nationality",
			MinLength:               global.IntPtr(2),
			MaxLength:               global.IntPtr(100),
			CanBeEmpty:              false,
			Regex:                   regexp.MustCompile(NameRegex),
			RegexFormatErrorMessage: &NameRegexMessage,
		},
		&StringFieldValidationData{
			FieldName:  "Bio",
			MaxLength:  global.IntPtr(5000),
			CanBeEmpty: true,
		},
		&StringFieldValidationData{
			FieldName:  "Website",
			MaxLength:  global.IntPtr(500),
			CanBeEmpty: true,
		},
	}
	CategoryValidation = []FieldValidationData{
		&StringFieldValidationData{
			FieldName:  "Name",
			MinLength:  global.IntPtr(2),
			MaxLength:  global.IntPtr(100),
			CanBeEmpty: false,
		},
		&StringFieldValidationData{
			FieldName:  "Description",
			MaxLength:  global.IntPtr(5000),
			CanBeEmpty: true,
		},
	}
	BookValidation = []FieldValidationData{
		&StringFieldValidationData{
			FieldName:  "Title",
			MinLength:  global.IntPtr(1),
			MaxLength:  global.IntPtr(100),
			CanBeEmpty: false,
		},
		&StringFieldValidationData{
			FieldName:  "Language",
			MinLength:  global.IntPtr(2),
			MaxLength:  global.IntPtr(100),
			CanBeEmpty: false,
		},
		&IntFieldValidationData{
			FieldName: "Year",
			MinValue:  global.IntPtr(0),
			MaxValue:  global.IntPtr(time.Now().Year()),
		},
		&IntFieldValidationData{
			FieldName: "Total Copies",
			MinValue:  global.IntPtr(0),
		},
	}
	ReviewValidation = []FieldValidationData{
		&StringFieldValidationData{
			FieldName:  "Content",
			MinLength:  global.IntPtr(1),
			MaxLength:  global.IntPtr(5000),
			CanBeEmpty: false,
		},
		&FloatFieldValidationData{
			FieldName: "Rating",
			MinValue:  global.FloatPtr(1),
			MaxValue:  global.FloatPtr(5),
		},
	}
)
