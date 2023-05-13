package slicew

import "fmt"

// ByPointer demonstrates passing slice by pointer.
func ByPointer(s *[]string) {
	(*s)[2] = "three"
	*s = append(*s, "byPointer")
	fmt.Printf("byPointer: %v\n", s)
}

// ByReference demonstrates passing slice by reference.
func ByReference(s []string) {
	s[3] = "four"
	s = append(s, "byReference")
	fmt.Printf("byReference: %v\n", s)
}

func SliceParameters() {
	// See ExampleSliceParameters
}
