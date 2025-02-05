package db

import (
	"fmt"
)

type ErrorType int

const (
	ValidationError ErrorType = iota
	NotFoundError
	InternalError
)

func (e ErrorType) String() string {
	switch e {
	case ValidationError:
		return "Validation Error"
	case NotFoundError:
		return "Not Found Error"
	case InternalError:
		return "Internal Error"
	default:
		return "Unknown Error"
	}
}

type DBError struct {
	Type ErrorType
	Err  string
}

func NewDBError(Type ErrorType, format string, args ...interface{}) *DBError {
	return &DBError{
		Type,
		fmt.Sprintf(format, args...),
	}
}

func (e DBError) Error() string {
	return e.Err
}
