package functionalw

import "fmt"

// NOTE:
// As is pointed out here: https://www.youtube.com/watch?v=rpB3P0QlvII
// the use of functional programming constructs can entail performance penalties.
// Always consider the reasons why you chose Go in the first place.

// Map function.
// From: https://stackoverflow.com/questions/71624828/is-there-a-way-to-map-an-array-of-objects-in-golang
func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

// MakeMul Example of Currying.
func MakeMul(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

func RunFunctionalWisdom(run bool) {

	if !run {
		return
	}

	data0 := []int{1, 2, 3, 4, 5}
	mul := 4
	mulFunc := MakeMul(mul)
	data1 := Map(data0, mulFunc)

	fmt.Printf("Original: %v\n", data0)
	fmt.Printf("Times %v: %v\n", mul, data1)
}
