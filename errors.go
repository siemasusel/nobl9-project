package stddevapi

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Parent  error
	Message string
}

func NewValidationError(parent error, msg string, args ...interface{}) error {
	return &ValidationError{
		Parent:  parent,
		Message: fmt.Sprintf(msg, args...),
	}
}

func (err ValidationError) Error() string {
	msg := fmt.Sprintf("Message=(%s)", err.Message)

	if err.Parent != nil {
		msg += fmt.Sprintf(" Parent=(%v)", err.Parent)
	}

	return msg
}

func (err *ValidationError) Is(target error) bool {
	var vErr *ValidationError

	return errors.As(err, &vErr)
}

func (err *ValidationError) Unwrap() error {
	return err.Parent
}

func IsValidationError(err error) bool {
	var vErr *ValidationError

	return errors.As(err, &vErr)
}
