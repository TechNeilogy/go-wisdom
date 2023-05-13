package structw

import (
	"fmt"
)

func Example_oop() {

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

	// output: Initial: 1
	// ByVal: 1
	// ByRef: 3
	// ByCopy (original): 3
	// ByCopy (copy): 4
	// ByCopyRef (original): 5
	// ByCopyRef (copy): 5
	// ByCopyVal (chained) (original): 5
	// ByCopyVal (chained) (copy): 8
	// ByCopyRef (chained) (original): 11
	// ByCopyRef (chained) (copy): 11
	// ByCopyRef (chained) (original): 11
	// ByCopyRef (chained) (copy): 14
}

func ExampleStringStructPenalty() {
	StringStructPenalty()
	// output: Outer Loops: 10000
	// Inner Loops: 10
	// Total Loops: 100000
	// Run:  1 2 3 4 5
	// Result of 5 Runs:
	// sumString > sumInt = true
	// meanString > meanInt = true
}

func ExampleIntPtrStructPenalty() {
	IntPtrStructPenalty()
	// output: Outer Loops: 10000
	// Inner Loops: 10
	// Total Loops: 100000
	// Run:  1 2 3 4 5
	// Result of 5 Runs:
	// sumIntPtr > sumInt = true
	// meanIntPtr > meanInt = true
}
