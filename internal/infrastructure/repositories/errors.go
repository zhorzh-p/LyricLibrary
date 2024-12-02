package repositories

import (
	"errors"
	"fmt"
)

var (
	ErrEntityNotFound = errors.New("entity not found")
)

type ErrRepositoryError struct {
	Message string
	Cause   error
}

func (e *ErrRepositoryError) Error() string {
	return fmt.Sprintf("Repository error.\nMessage: %s.\nCause: %s", e.Message, e.Cause.Error())
}

func NewErrRepositoryError(message string, cause error) error {
	return &ErrRepositoryError{message, cause}
}
