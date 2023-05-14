package slicew

import "fmt"

func ExampleSliceParameters() {
	s := []string{"1", "2", "3", "4", "5"}
	ByPointer(&s)
	ByReference(s)
	fmt.Printf("Run: %v\n", s)

	// output byPointer: &[1 2 three 4 5 byPointer]
	// byReference: [1 2 three four 5 byPointer byReference]
	// Run: [1 2 three four 5 byPointer]
}
