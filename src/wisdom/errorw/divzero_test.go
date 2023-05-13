package errorw

import "fmt"

func ExampleDivZero_Error() {

	// div is a custom division function.
	div := func(a int, b int) (int, error) {
		if b == 0 {
			return 0, DivZero{a, b}
		}
		return a / b, nil
	}

	fmt.Println(div(10, 2))
	fmt.Println(div(0, 10))
	fmt.Println(div(10, 0))
	fmt.Println(div(0, 0))

	// output:5 <nil>
	//0 <nil>
	//0 division by zero: 10 / 0
	//0 division by zero: 0 / 0
}
