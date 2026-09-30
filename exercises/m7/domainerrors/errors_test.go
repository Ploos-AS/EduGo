package domainerrors

import (
	"errors"
	"testing"
)

func TestLookupNotFound(t *testing.T) {
	err := Lookup(42)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestLookupValidation(t *testing.T) {
	err := Lookup(0)
	var validation *ValidationError
	if !errors.As(err, &validation) {
		t.Fatalf("expected ValidationError, got %v", err)
	}
	if validation.Field != "id" {
		t.Fatalf("field = %q", validation.Field)
	}
}
