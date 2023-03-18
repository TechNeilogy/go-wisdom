package functionalw

import "fmt"

func ExampleMakeMul() {

	data0 := []int{0, 1, 2, 3, 4}
	mul := 4
	mulFunc := MakeMul(mul)
	data1 := Map(data0, mulFunc)
	fmt.Printf("%v ", data1)

	// output: [0 4 8 12 16]
}

func ExampleMap() {

	fmt.Printf(
		"%v",
		Map(
			[]int{0, 1, 2, 3, 4},
			func(i int) string { return fmt.Sprintf("x%v", i*2) },
		),
	)

	// output: [x0 x2 x4 x6 x8]
}

func ExampleMyStream_Filter() {

	stream0 := MyStream[int]{
		[]int{0, 1, 2, 3, 4},
	}

	stream1 := stream0.
		Filter(func(x int) bool { return x > 0 }).
		Filter(func(x int) bool { return x < 4 })

	fmt.Printf("%v ", stream1.data)

	// output: [1 2 3]
}
