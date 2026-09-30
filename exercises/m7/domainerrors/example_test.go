package domainerrors_test

import (
	"errors"
	"fmt"

	"github.com/Ploos-AS/EduGo/exercises/m7/domainerrors"
)

func ExampleLookup() {
	err := domainerrors.Lookup(7)
	fmt.Println(errors.Is(err, domainerrors.ErrNotFound))
	// Output:
	// true
}
