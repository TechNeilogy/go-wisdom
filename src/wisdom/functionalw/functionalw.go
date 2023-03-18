package functionalw

// NOTE:
// As is pointed out here: https://www.youtube.com/watch?v=rpB3P0QlvII
// the use of functional programming constructs can entail performance penalties.
// Always consider the reasons why you chose Go in the first place.

// Map maps a function onto a slice.
// From: https://stackoverflow.com/questions/71624828/is-there-a-way-to-map-an-array-of-objects-in-golang
func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

// MyStream is a simple mock stream for use in Filter.
type MyStream[T any] struct {
	data []T
}

// Filter filters a stream based on a predicate.
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

// MakeMul shows an xample of Currying.
func MakeMul(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}
