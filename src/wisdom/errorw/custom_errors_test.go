package errorw

import "fmt"

func ExampleDivZero_Error() {

	div := func(a int, b int) (int, error) {
		if b == 0 {
			return 0, DivZero{a, b}
		}
		return a / b, nil
	}

	_, err0 := div(10, 2)
	_, err1 := div(0, 10)
	_, err2 := div(10, 0)
	_, err3 := div(0, 0)

	fmt.Print(
		err0 == nil,
		err1 == nil,
		err2 == nil,
		err3 == nil,
	)

	// output: true true false false
}
