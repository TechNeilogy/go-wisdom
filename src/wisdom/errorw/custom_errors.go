package errorw

import (
	"fmt"
)

// DivZero is a custom divide by zero error.
type DivZero struct {
	v0 int
	v1 int
}

// Error implements the error interface for DivZero.
func (e DivZero) Error() string {
	return fmt.Sprintf("division by zero: %v / %v", e.v0, e.v1)
}
