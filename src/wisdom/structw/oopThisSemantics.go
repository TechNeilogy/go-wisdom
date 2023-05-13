package structw

import (
	"fmt"
)

type oop struct {
	a int
}

func (o oop) String() string {
	return fmt.Sprintf("%v", o.a)
}

// Pass struct to copy.
func (o oop) byval(a int) {
	o.a = a
}

// Pass struct pointer to modify.
func (o *oop) byref(a int) {
	o.a = a
}

// Return a copy.
func (o oop) bycopy(a int) oop {
	o.a = a
	return o
}

// Return a copy.
func (o oop) bycopyval(a int) oop {
	o.a = a
	return o
}

// Return origonal.
func (o *oop) bycopyref(a int) *oop {
	o.a = a
	return o
}
