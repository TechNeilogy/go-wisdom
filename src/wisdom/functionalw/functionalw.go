package functionalw

import (
	"fmt"
	"github.com/TechNeilogy/go-wisdom/src/util"
)

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

type MyStream[T any] struct {
	data []T
}

func (s *MyStream[T]) Filter(f func(T) bool) *MyStream[T] {
	var rtn []T
	for i := range s.data {
		if f(s.data[i]) {
			rtn = append(rtn, s.data[i])
		}
	}
	return &MyStream[T]{
		rtn,
	}
}

// MakeMul Example of Currying.
func MakeMul(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

func runMappingAndFiltering() {

	data0 := []int{1, 2, 3, 4, 5}
	mul := 4
	mulFunc := MakeMul(mul)
	data1 := Map(data0, mulFunc)

	stream1 := MyStream[int]{
		data1,
	}

	// Note: Much slower than just a single loop.
	stream2 := stream1.
		Filter(func(x int) bool { return x > 10 }).
		Filter(func(x int) bool { return x < 20 })

	util.PrintHeader("Functional Mapping and Filtering")

	fmt.Printf("Original: %v\n", data0)
	fmt.Printf("Map Times %v: %v\n", mul, data1)
	fmt.Printf("Filter 10 < x < 20: %v\n", stream2.data)
}

func RunFunctionalWisdom(run bool) {

	if !run {
		return
	}

	runMappingAndFiltering()
}
