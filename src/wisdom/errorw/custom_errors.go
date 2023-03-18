package errorw

import (
	"fmt"
)

type DivZero struct {
	v0 int
	v1 int
}

func (e DivZero) Error() string {
	return fmt.Sprintf("division by zero: %v / %v", e.v0, e.v1)
}
