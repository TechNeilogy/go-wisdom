package oopthissemantics

import (
	"fmt"
	"github.com/TechNeilogy/go-wisdom/src/util"
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

func Run() {

	util.PrintHeader("OOP 'this' Passing Semantics")

	o := oop{
		a: 1,
	}

	fmt.Printf("Initial: %v\n", o)

	o.byval(2)

	fmt.Printf("ByVal: %v\n", o)

	o.byref(3)

	fmt.Printf("ByRef: %v\n", o)

	o0 := o.bycopy(4)

	fmt.Printf("ByCopy (original): %v\n", o)
	fmt.Printf("ByCopy (copy): %v\n", o0)

	o1 := o.bycopyref(5)

	fmt.Printf("ByCopyRef (original): %v\n", o)
	fmt.Printf("ByCopyRef (copy): %v\n", o1)

	// Note "dot" break syntax.
	o3 := o.
		bycopyval(6).
		bycopyval(7).
		bycopyval(8)

	fmt.Printf("ByCopyVal (chained) (original): %v\n", o)
	fmt.Printf("ByCopyVal (chained) (copy): %v\n", o3)

	// Note "dot" break syntax.
	o4 := o.
		bycopyref(9).
		bycopyref(10).
		bycopyref(11)

	fmt.Printf("ByCopyRef (chained) (original): %v\n", o)
	fmt.Printf("ByCopyRef (chained) (copy): %v\n", o4)

	// Preserve original, but minimize copying.
	o5 := o
	o6 := o5.
		bycopyref(12).
		bycopyref(13).
		bycopyref(14)

	fmt.Printf("ByCopyRef (chained) (original): %v\n", o)
	fmt.Printf("ByCopyRef (chained) (copy): %v\n", o6)
}
