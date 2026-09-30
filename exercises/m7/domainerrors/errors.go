// Package domainerrors demonstrates errors that callers can classify
// without depending on human-readable error strings.
package domainerrors

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid field %q", e.Field)
}

func Lookup(id int) error {
	if id <= 0 {
		return &ValidationError{Field: "id"}
	}
	return fmt.Errorf("lookup %d: %w", id, ErrNotFound)
}
