package mapReduce

import (
	"context"
	"fmt"
)

// Fib is a naive Fibonacci implementation,
// to use up CPU cycles.
func Fib(n int) int {
	if n > 42 {
		panic(fmt.Sprintf("did you really mean to compute fib(%v)?", n))
	}
	if n < 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func ExampleMapReduce() {
	ctx := context.TODO()

	il := MakeIterList(
		[][]int{
			{0, 1, 2, 3, 4, 41},
			{5, 6, 7, 8, 9, 41},
			{0, 1, 2, 3, 4, 41},
			{5, 6, 7, 8, 9, 41},
			{0, 1, 2, 3, 4, 41},
			{5, 6, 7, 8, 9, 41},
			{0, 1, 2, 3, 4, 41},
			{5, 6, 7, 8, 9, 41},
		},
	)

	sum := MapReduce[[]int](
		ctx,
		// Map.
		func(
			ctx context.Context,
			l []int,
		) []int {
			var l0 []int
			for _, v := range l {
				l0 = append(l0, Fib(v))
			}
			return l0
		},
		// Reduce.
		func(
			ctx context.Context,
			u []int,
			acc int,
		) int {
			for _, i := range u {
				acc += i
			}
			return acc
		},
		&il,
		true,
	)

	fmt.Printf("%v", sum)

	//output: 2143314940
}
